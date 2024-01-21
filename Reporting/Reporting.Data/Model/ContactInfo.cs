namespace Reporting.Data.Model;

public class ContactInfo
{
    public Guid Id { get; set; }
    
    public Guid HotelId { get; set; }
    
    public string Name { get; set; }
    
    public string Phone { get; set; }
    
    public string Email { get; set; }
}