using System;
using System.IO; 

namespace hello
{
    class Program
    {
        static void Main(string[] args)
        {
            if(args == null || args.Length < 1){
                Console.WriteLine("Please provide a file path with the hello text...");
            }

            string path = args[0];
            
            if(!File.Exists(path)){
                throw new FileNotFoundException($"Could not find file at {path}");
            }

            string helloText = File.ReadAllText(path);
            string updatedHello = updateHelloText(helloText);

            Console.WriteLine(updatedHello);
        }

        private static string updateHelloText(string text){
            string updatedText = text.Replace("shell command", "csharp program");
            return updatedText;
        }
    }
}
