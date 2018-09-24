using System;
using System.Collections.Generic;
using System.Text;

class Braces
{
    static void Main()
    {
        Stack<char> bracesStack = new Stack<char>();

        int n = int.Parse(Console.ReadLine());

        StringBuilder output = new StringBuilder();
        for (int i = 0; i < n; i++)
        {
            string line = Console.ReadLine();
            output.AppendLine(AreBracesValid(line) ? "YES" : "NO");
        }

        Console.WriteLine(output.ToString());
    }

    static bool AreBracesValid(string line)
    {
        Stack<char> bracesStack = new Stack<char>();

        foreach (char ch in line)
        {
            if (IsOpeningBrace(ch))
            {
                bracesStack.Push(ch);
            }
            else if (IsClosingBrace(ch))
            {
                char topChar = bracesStack.Pop();
                if (!BracesMatch(topChar, ch))
                {
                    return false;
                }
            }
        }

        return bracesStack.Count == 0;
    }

    static bool IsOpeningBrace(char ch)
    {
        return ch == '{' || ch == '(' || ch == '[';
    }

    static bool IsClosingBrace(char ch)
    {
        return ch == '}' || ch == ')' || ch == ']';
    }

    static bool BracesMatch(char openingBrace, char closingBrace)
    {
        return (openingBrace == '(' && closingBrace == ')') ||
               (openingBrace == '{' && closingBrace == '}') ||
               (openingBrace == '[' && closingBrace == ']');
    }
}