using Reporting.Data.DTOs;
using Reporting.Data.Model;

namespace ReportingService.Services;

public interface IReportService
{
    Task<Report> CreateReportAsync(ReportCreateDto reportCreateDto);
    
    Task<Report> UpdateReportAsync(Guid id, ReportUpdateDto reportUpdateDto);
}