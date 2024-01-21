using Microsoft.EntityFrameworkCore;
using Reporting.Data;
using Reporting.Data.DTOs;
using Reporting.Data.Model;
using Reporting.Data.Utils;

namespace ReportingService.Services;

public class ReportService : IReportService
{
    private readonly ApplicationDbContext _context;

    public ReportService(ApplicationDbContext context)
    {
        _context = context;
    }

    public async Task<Report> CreateReportAsync(ReportCreateDto reportCreateDto)
    {
        Report report = new();
        report.RequestedDate = DateTime.SpecifyKind(DateTime.Now, DateTimeKind.Utc);
        report.Status = ReportStatus.Pending;
        report.City = reportCreateDto.City;
        report.Country = reportCreateDto.Country;
        report.HotelCount = 0;
        report.PhoneNumberCount = 0;

        await _context.Reports.AddAsync(report);
        await _context.SaveChangesAsync();

        return report;
    }

    public async Task<Report> UpdateReportAsync(Guid id, ReportUpdateDto reportUpdateDto)
    {
        var report = await _context.Reports.FirstOrDefaultAsync(x => x.Id == id);

        if (report is null)
        {
            throw new InvalidOperationException("Report not found");
        }

        report.Status = reportUpdateDto.Status;
        report.HotelCount = reportUpdateDto.HotelCount;
        report.PhoneNumberCount = reportUpdateDto.PhoneNumberCount;

        _context.Reports.Update(report);

        await _context.SaveChangesAsync();

        return report;
    }
}