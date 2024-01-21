using System.Text;
using System.Text.Json;

namespace ReportingService.Utils;

public class APIResponse<T>
{
    public bool Status { get; set; }
    public string Message { get; set; } = string.Empty;
    public T Data { get; set; }
    public List<Error> Errors { get; set; } = new();
    
    public void AddErrorAsync(string code, string message)
    {
        this.Errors.Add(new Error
        {
            Code = code,
            Message = message
        });
    }

    public string GetJson()
    {
        StringBuilder sb = new();

        sb.Append("{");
        sb.Append($"\"status\":{Status.ToString().ToLower()},");
        sb.Append($"\"message\":\"{Message}\",");
        sb.Append($"\"data\":{JsonSerializer.Serialize<T>(Data)},");
        sb.Append($"\"errors\":[{GetErrorJson()}]");
        sb.Append("}");
        
        return sb.ToString();
    }
    
    public string GetErrorJson()
    {
        StringBuilder sb = new();

        foreach (var error in this.Errors)
        {
            sb.Append(error.GetJson());

            if (error != this.Errors.Last())
            {
                sb.Append(",");
            }
        }

        return sb.ToString();
    }
}

public class Error
{
    public string Code { get; set; } = string.Empty;
    public string Message { get; set; } = string.Empty;
    
    // we will create a method to get the json
    public string GetJson()
    {
        StringBuilder sb = new();

        sb.Append("{");
        sb.Append($"\"Code\":\"{Code}\",");
        sb.Append($"\"Message\":\"{Message}\"");
        sb.Append("}");

        return sb.ToString();
    }
}