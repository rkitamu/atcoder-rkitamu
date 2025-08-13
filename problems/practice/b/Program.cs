using AtCoder;
using MathNet;

class Program
{
    static void Main()
    {
        var abc = Console.ReadLine();
        var cnt = 0;
        for (int i = 0; i < abc.Length; i++)
        {
            if (abc[i] == '1')
            {
                cnt++;
            }
        }
        Console.WriteLine(cnt);
    }
}
