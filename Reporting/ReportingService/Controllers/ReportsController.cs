using Microsoft.AspNetCore.Mvc;
using Reporting.Data.DTOs;

namespace ReportingService.Controllers;

[ApiController]
[Produces("application/json")]
[Route("[controller]")]
public class ReportsController : ControllerBase
{
    [HttpPost]
    public async Task<IActionResult> Post(ReportCreateDto reportCreateDto)
    {
        return Accepted();
    }
}