using System.Text;
using RabbitMQ.Client;

namespace ReportingService.Messaging;

public class RabbitMQPublisher : IMessagePublisher
{
    public RabbitMQPublisher()
    {
    }

    public async Task PublishReportRequestAsync(string reportRequestJson)
    {
        throw new NotImplementedException();
    }
}