using System.Text;
using RabbitMQ.Client;

namespace ReportingService.Messaging;

public class RabbitMQPublisher : IMessagePublisher
{
    private readonly IConnection _connection;
    private readonly IModel _channel;

    public RabbitMQPublisher()
    {
        ConnectionFactory factory = new();
        factory.HostName = "localhost";
        factory.Port = 5672;
        factory.UserName = "guest";
        factory.Password = "guest";

        _connection = factory.CreateConnection();
        _channel = _connection.CreateModel();

        _channel.QueueDeclare("report-requests",
            false,
            false,
            false,
            null);
    }

    public async Task PublishReportRequestAsync(string reportRequestJson)
    {
        byte[] body = Encoding.UTF8.GetBytes(reportRequestJson);

        _channel.BasicPublish("", "report-requests", null, body);

        await Task.CompletedTask;
    }
    
    ~RabbitMQPublisher()
    {
        _channel.Close();
        _connection.Close();
    }
}