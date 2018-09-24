using System;

class Music
{
    static void Main()
    {
        int N = int.Parse(Console.ReadLine());
        int K = int.Parse(Console.ReadLine());
        int L = int.Parse(Console.ReadLine());

        long res = 1;
        for (int i = 0; i < L; i++)
        {
            if (i < N)
            {
                res *= N - i;
            }
            else
            {
                res *= Math.Max(N - i, N - K);
            }

            res %= 1000000007;
        }

        Console.WriteLine(res);
    }
}
