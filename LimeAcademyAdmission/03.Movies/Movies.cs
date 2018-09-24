using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.IO;
using System.Net;
using System.Web;
using System.Linq;

class Movies
{
    static void Main()
    {
        string substr = Console.ReadLine();
        List<string> titles = new List<string>();
        int currentPage = 0;
        int totalPages = 0;

        do
        {
            currentPage++;

            string html = string.Empty;
            string url = 
                String.Format("https://jsonmock.hackerrank.com/api/movies/search/" +
                              "?Title={0}&page={1}", substr, currentPage);

            HttpWebRequest request = (HttpWebRequest)WebRequest.Create(url);
            using (HttpWebResponse response = (HttpWebResponse)request.GetResponse())
            using (Stream stream = response.GetResponseStream())
            using (StreamReader reader = new StreamReader(stream))
            {
                html = reader.ReadToEnd();
            }

            dynamic json = JsonConvert.DeserializeObject(html);
            totalPages = json["total_pages"];

            for (int i = 0; i < json["data"].Count; i++)
            {
                titles.Add(json["data"][i]["Title"].ToString());
            }

        } while (currentPage <= totalPages);

        titles.Sort();
        Console.WriteLine(string.Join("" + Environment.NewLine, titles));
    }
}
