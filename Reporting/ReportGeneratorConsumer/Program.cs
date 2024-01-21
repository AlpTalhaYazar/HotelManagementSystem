using System.Text;
using System.Text.Json;
using Microsoft.EntityFrameworkCore;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using Reporting.Data;
using Reporting.Data.Model;
using Reporting.Data.Utils;

namespace ReportGeneratorConsumer;

public class Program
{
    static void Main()
    {
        ConnectionFactory factory = new();
        factory.HostName = "localhost";
        factory.Port = 5672;
        factory.UserName = "guest";
        factory.Password = "guest";

        using var connection = factory.CreateConnection();
        using var channel = connection.CreateModel();

        channel.QueueDeclare(queue: "report-requests",
            durable: false,
            exclusive: false,
            autoDelete: false,
            arguments: null);

        var consumer = new EventingBasicConsumer(channel);
        consumer.Received += (model, ea) =>
        {
            var body = ea.Body.ToArray();
            var message = Encoding.UTF8.GetString(body);
            Console.WriteLine("Received message: " + message);

            // Process the report request
            ProcessReport(message);
        };

        channel.BasicConsume(queue: "report-requests",
            autoAck: true,
            consumer: consumer);

        Console.WriteLine("Press [enter] to exit.");
        Console.ReadLine();
    }

    private static async void ProcessReport(string message)
    {
        Console.WriteLine("Processing report");
        
        DbContextOptions<ApplicationDbContext> dbContextOptions = new DbContextOptionsBuilder<ApplicationDbContext>()
            .UseNpgsql("Host=localhost;Port=5432;Database=hotel_management_system;Username=postgres").Options;
        
        using (ApplicationDbContext context = new(dbContextOptions))
        {
            Report report = JsonSerializer.Deserialize<Report>(message);
        
            report.Status = ReportStatus.Processing;
            context.Reports.Update(report);
            await context.SaveChangesAsync();
            
            await Task.Delay(10000); // Simulate processing time
            
            var hotels = await context.Hotels
                .Where(x => x.City == report.City && x.Country == report.Country)
                .ToListAsync();
            
            var hotelIds = hotels.Select(x => x.Id)
                .ToList();
            
            var contactInfos = await context.ContactInfos
                .Where(x => hotelIds.Contains(x.HotelId))
                .ToListAsync();
            
            report.HotelCount = hotels.Count;
            report.PhoneNumberCount = contactInfos.Count;
            report.Status = ReportStatus.Completed;
            
            context.Reports.Update(report);
            await context.SaveChangesAsync();
        }

        Console.WriteLine("Report processed and completed");
    }
}