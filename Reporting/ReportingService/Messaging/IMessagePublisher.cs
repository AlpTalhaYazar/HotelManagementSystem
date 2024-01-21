namespace ReportingService.Messaging;

public interface IMessagePublisher
{
    Task PublishReportRequestAsync(string reportRequestJson);
}