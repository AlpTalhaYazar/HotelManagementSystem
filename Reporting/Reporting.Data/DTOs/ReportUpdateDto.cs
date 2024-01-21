using Reporting.Data.Utils;

namespace Reporting.Data.DTOs;

public class ReportUpdateDto
{
    public ReportStatus Status { get; set; }
    public int HotelCount { get; set; }
    public int PhoneNumberCount { get; set; }
}