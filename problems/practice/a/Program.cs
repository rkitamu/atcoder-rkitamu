class Program
{
    static void Main()
    {
        int a = int.Parse(Console.ReadLine()!);
        var (b, c) = Console.ReadLine()!
            .Split()
            .Select(int.Parse)
            .ToArray() switch
        { var arr => (arr[0], arr[1]) };
        string s = Console.ReadLine()!;
        Console.WriteLine($"{a + b + c} {s}");
    }
}
