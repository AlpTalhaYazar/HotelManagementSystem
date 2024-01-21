using System.Text.Json;
using Microsoft.AspNetCore.Mvc;
using Reporting.Data.DTOs;
using Reporting.Data.Model;
using ReportingService.Messaging;
using ReportingService.Services;
using ReportingService.Utils;

namespace ReportingService.Controllers;

[ApiController]
[Produces("application/json")]
[Route("[controller]")]
public class ReportsController : ControllerBase
{
    private readonly IReportService _reportService;

    public ReportsController(IMessagePublisher messagePublisher, IReportService reportService)
    {
        _reportService = reportService;
    }

    [HttpPost]
    public async Task<IActionResult> Post(ReportCreateDto reportCreateDto)
    {
        Report report = await _reportService.CreateReportAsync(reportCreateDto);

        APIResponse<Report> apiResponse = new();
        apiResponse.Status = true;
        apiResponse.Message = "Report request has been sent successfully";
        apiResponse.Data = report;

        return Accepted(apiResponse);
    }
}