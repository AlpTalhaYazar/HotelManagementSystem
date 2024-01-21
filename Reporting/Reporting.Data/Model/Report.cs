using Reporting.Data.Utils;

namespace Reporting.Data.Model;

public class Report
{
    public Guid Id { get; set; }
    public DateTime RequestedDate { get; set; }
    public ReportStatus Status { get; set; }
    public string City { get; set; }
    public string Country { get; set; }
    public int HotelCount { get; set; }
    public int PhoneNumberCount { get; set; }
}