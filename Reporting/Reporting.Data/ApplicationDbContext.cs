using Microsoft.EntityFrameworkCore;
using Reporting.Data.Model;

namespace Reporting.Data;

public class ApplicationDbContext : DbContext
{
    public ApplicationDbContext(DbContextOptions<ApplicationDbContext> options) : base(options)
    {
    }

    public DbSet<Report> Reports { get; set; }
    public DbSet<Hotel> Hotels { get; set; }
    public DbSet<ContactInfo> ContactInfos { get; set; }
}