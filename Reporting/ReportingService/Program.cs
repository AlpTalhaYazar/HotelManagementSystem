using Microsoft.EntityFrameworkCore;
using Reporting.Data;
using ReportingService.Utils;

var builder = WebApplication.CreateBuilder(args);

var databaseConfigurations = builder.Configuration.GetSection("Database").Get<List<DatabaseConfiguration>>();

var defaultDatabaseConfiguration = databaseConfigurations.FirstOrDefault(x => x.Name == "Default");

if (defaultDatabaseConfiguration is null)
{
    throw new Exception("Default database configuration not found");
}

builder.Services.AddDbContext<ApplicationDbContext>(options =>
    options.UseNpgsql(defaultDatabaseConfiguration.ConnectionString));

// Add services to the container.
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

app.UseRouting();

app.UseEndpoints(endpoints =>
{
    endpoints.MapControllers();
});

app.Run();