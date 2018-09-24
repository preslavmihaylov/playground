using System;

class FizzBuzz
{
    static void Main()
    {
        int n = int.Parse(Console.ReadLine());

        for (int i = 1; i <= n; i++)
        {
            if (i % 3 == 0 && i % 5 == 0)
            {
                Console.WriteLine("{0} -> FizzBuzz", i);
            }
            else if (i % 3 == 0)
            {
                Console.WriteLine("{0} -> Fizz", i);
            }
            else if (i % 5 == 0)
            {
                Console.WriteLine("{0} -> Buzz", i);
            }
            else
            {
                Console.WriteLine("{0} -> {0}", i);
            }
        }
    }
}