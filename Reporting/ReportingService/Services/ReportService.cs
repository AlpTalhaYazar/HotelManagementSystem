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
        throw new NotImplementedException();
    }

    public async Task<Report> UpdateReportAsync(Guid id, ReportUpdateDto reportUpdateDto)
    {
        throw new NotImplementedException();
    }
}