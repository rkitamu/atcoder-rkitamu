using AtCoder;
using MathNet;

class Program
{
    static void Main()
    {
        var n = StdReader.ReadSingle<int>();
        var a = StdReader.ReadSingle<int>();
        var b = StdReader.ReadSingle<int>();

        var ans = 0;
        for (int i = 1; i <= n; i++)
        {
            var s = i.ToString();
            var sum = 0;
            for (int j = 0; j < s.Length; j++)
            {
                // convert s[j] to int
                sum += s[j] - '0';
            }
            if (a <= sum && sum <= b)
            {
                ans += i;
            }
        }
        StdWriter.PrintLine(ans);
    }
}

public static class StdReader
{
    // 高速な標準入力リーダー。空白区切りのトークン/行/グリッド/行列を読む。
    private static TextReader _reader = Console.In;

    // 現在行の未消費トークン
    private static readonly Queue<string> _tokens = new();

    /// <summary>入力元を差し替える（既存トークンは破棄）。既定は Console.In。</summary>
    public static void SetInput(TextReader reader)
    {
        _reader = reader ?? throw new ArgumentNullException(nameof(reader));
        _tokens.Clear();
    }

    /// <summary>
    /// 次のトークンを T に変換して返す（空白区切り）。
    /// </summary>
    /// <typeparam name="T">string, int, long, double, char のみ対応。</typeparam>
    /// <exception cref="InvalidOperationException">入力が尽きた。</exception>
    public static T ReadSingle<T>()
    {
        EnsureTokenAvailable();
        return Converter<T>.Parse(_tokens.Dequeue());
    }

    /// <summary>
    /// 次の物理1行を空白で分割し、各要素を T に変換して返す。
    /// 行途中で止めたい場合は事前に <see cref="DiscardRestOfLine"/> を呼ぶ。
    /// </summary>
    /// <typeparam name="T">string, int, long, double, char のみ対応。</typeparam>
    /// <exception cref="InvalidOperationException">入力が尽きた。</exception>
    public static T[] ReadMultiple<T>()
    {
        var line = ReadLineOrThrow();
        var parts = line.Split((char[])null, StringSplitOptions.RemoveEmptyEntries);
        var result = new T[parts.Length];
        for (int i = 0; i < parts.Length; i++)
            result[i] = Converter<T>.Parse(parts[i]);
        return result;
    }

    /// <summary>
    /// 連続文字列のグリッド h×w を読み取り、左右上下に offset の余白を付けて返す。
    /// 余白は '\0' で埋める。各行の長さが w でないときは例外。
    /// </summary>
    /// <exception cref="ArgumentOutOfRangeException">h, w, offset が負。</exception>
    /// <exception cref="InvalidOperationException">入力が尽きた。</exception>
    /// <exception cref="FormatException">行の長さが一致しない。</exception>
    public static char[][] ReadGrid(int h, int w, int offset = 0)
    {
        if (h < 0) throw new ArgumentOutOfRangeException(nameof(h));
        if (w < 0) throw new ArgumentOutOfRangeException(nameof(w));
        if (offset < 0) throw new ArgumentOutOfRangeException(nameof(offset));

        var totalRows = h + 2 * offset;
        var totalCols = w + 2 * offset;
        var grid = new char[totalRows][];

        // 上余白
        for (int i = 0; i < offset; i++)
            grid[i] = new char[totalCols];

        // 本体（左右に余白）
        for (int i = offset; i < offset + h; i++)
        {
            var line = ReadLineOrThrow();
            if (line.Length != w)
                throw new FormatException($"Line {i - offset}: expected {w}, got {line.Length}.");

            var row = new char[totalCols];
            for (int j = 0; j < w; j++)
                row[offset + j] = line[j];
            grid[i] = row;
        }

        // 下余白
        for (int i = offset + h; i < totalRows; i++)
            grid[i] = new char[totalCols];

        return grid;
    }

    /// <summary>
    /// 空白区切りの数表 h×w を読み取り、左右上下に offset の余白を付けて返す。
    /// 余白は default(T)。
    /// </summary>
    /// <typeparam name="T">string, int, long, double, char のみ対応。</typeparam>
    /// <exception cref="ArgumentOutOfRangeException">h, w, offset が負。</exception>
    /// <exception cref="InvalidOperationException">入力が尽きた。</exception>
    public static T[][] ReadMatrix<T>(int h, int w, int offset = 0)
    {
        if (h < 0) throw new ArgumentOutOfRangeException(nameof(h));
        if (w < 0) throw new ArgumentOutOfRangeException(nameof(w));
        if (offset < 0) throw new ArgumentOutOfRangeException(nameof(offset));

        var totalRows = h + 2 * offset;
        var totalCols = w + 2 * offset;
        var a = new T[totalRows][];

        // 上余白
        for (int i = 0; i < offset; i++)
            a[i] = new T[totalCols];

        // 本体（左右に余白）
        for (int i = offset; i < offset + h; i++)
        {
            var row = new T[totalCols];
            for (int j = 0; j < w; j++)
                row[offset + j] = ReadSingle<T>();
            a[i] = row;
        }

        // 下余白
        for (int i = offset + h; i < totalRows; i++)
            a[i] = new T[totalCols];

        return a;
    }

    /// <summary>現在行の未消費トークンを破棄する（次回は次行の先頭から）。</summary>
    public static void DiscardRestOfLine() => _tokens.Clear();

    // ===== 内部実装（行読みの一元化） =====

    /// <summary>必要なら次行を読み、トークンを補充。</summary>
    private static void EnsureTokenAvailable()
    {
        while (_tokens.Count == 0)
        {
            var line = ReadLineOrThrow();
            foreach (var token in line.Split((char[])null, StringSplitOptions.RemoveEmptyEntries))
                _tokens.Enqueue(token);
        }
    }

    /// <summary>1行読み取り。EOF なら例外。</summary>
    private static string ReadLineOrThrow()
    {
        var line = _reader.ReadLine();
        if (line == null) throw new InvalidOperationException("No more input.");
        return line;
    }

    /// <summary>型ごとのパーサーを静的にキャッシュ。</summary>
    private static class Converter<T>
    {
        public static readonly Func<string, T> Parse = Create();

        private static Func<string, T> Create()
        {
            if (typeof(T) == typeof(string))
                return s => (T)(object)s;

            if (typeof(T) == typeof(int))
                return s => (T)(object)int.Parse(
                    s,
                    System.Globalization.NumberStyles.Integer,
                    System.Globalization.CultureInfo.InvariantCulture);

            if (typeof(T) == typeof(long))
                return s => (T)(object)long.Parse(
                    s,
                    System.Globalization.NumberStyles.Integer,
                    System.Globalization.CultureInfo.InvariantCulture);

            if (typeof(T) == typeof(double))
                return s => (T)(object)double.Parse(
                    s,
                    System.Globalization.NumberStyles.Float | System.Globalization.NumberStyles.AllowThousands,
                    System.Globalization.CultureInfo.InvariantCulture);

            if (typeof(T) == typeof(char))
                return s =>
                {
                    if (s.Length != 1)
                        throw new FormatException($"Token \"{s}\" is not a single character.");
                    return (T)(object)s[0];
                };

            return _ => throw new NotSupportedException($"Type {typeof(T)} is not supported.");
        }
    }
}

public static class StdWriter
{
    // 高速な標準出力。改行は "\n"、プロセス終了時に自動 Flush。
    static StdWriter()
    {
        AppDomain.CurrentDomain.ProcessExit += (_, __) =>
        {
            try { _writer.Flush(); } catch { /* 無視 */ }
        };
    }

    // バッファ付き Writer（AutoFlush=false）
    private static readonly StreamWriter _writer = new StreamWriter(Console.OpenStandardOutput())
    {
        AutoFlush = false,
        NewLine = "\n"
    };

    /// <summary>"Yes\n" を出力。</summary>
    public static void Yes() => _writer.Write("Yes\n");

    /// <summary>"No\n" を出力。</summary>
    public static void No() => _writer.Write("No\n");

    /// <summary>単一値を出力（改行付き）。</summary>
    public static void PrintLine<T>(T value)
    {
        if (value is IFormattable f)
            _writer.Write(f.ToString(null, System.Globalization.CultureInfo.InvariantCulture));
        else
            _writer.Write(value?.ToString() ?? "");
        _writer.Write('\n');
    }

    /// <summary>double を固定小数で出力（指数なし・末尾ゼロ除去・改行付き）。</summary>
    public static void PrintLine(double value, int digits)
    {
        _writer.Write(FormatDouble(value, digits));
        _writer.Write('\n');
    }

    /// <summary>1次元配列を空白区切りで出力（改行付き）。</summary>
    public static void PrintLine<T>(T[] a)
    {
        for (int i = 0; i < a.Length; i++)
        {
            if (i > 0) _writer.Write(' ');
            WriteOne(a[i]);
        }
        _writer.Write('\n');
    }

    /// <summary>double 配列を空白区切りで出力（指数なし・末尾ゼロ除去・改行付き）。</summary>
    public static void PrintLine(double[] a, int digits)
    {
        for (int i = 0; i < a.Length; i++)
        {
            if (i > 0) _writer.Write(' ');
            _writer.Write(FormatDouble(a[i], digits));
        }
        _writer.Write('\n');
    }

    /// <summary>2次元配列を行ごとに出力。</summary>
    public static void PrintLines<T>(T[][] mat)
    {
        for (int r = 0; r < mat.Length; r++)
            PrintLine(mat[r]);
    }

    /// <summary>double の2次元配列を行ごとに出力（指数なし・末尾ゼロ除去）。</summary>
    public static void PrintLines(double[][] mat, int digits)
    {
        for (int r = 0; r < mat.Length; r++)
            PrintLine(mat[r], digits);
    }

    // 要素1つの書き込み（double は既定15桁で末尾ゼロ除去）
    private static void WriteOne<T>(T v)
    {
        if (v is double d)
            _writer.Write(FormatDouble(d, 15));
        else if (v is IFormattable f)
            _writer.Write(f.ToString(null, System.Globalization.CultureInfo.InvariantCulture));
        else
            _writer.Write(v?.ToString() ?? "");
    }

    // double を指数なし・末尾ゼロ除去で文字列化
    private static string FormatDouble(double value, int digits)
    {
        if (double.IsNaN(value) || double.IsInfinity(value))
            return value.ToString(System.Globalization.CultureInfo.InvariantCulture);

        string s = value.ToString($"F{ClampDigits(digits)}", System.Globalization.CultureInfo.InvariantCulture);
        s = s.TrimEnd('0').TrimEnd('.');
        return s;
    }

    // 桁数の範囲を制限
    private static int ClampDigits(int d) => d < 0 ? 0 : (d > 99 ? 99 : d);
}

public static class Algorithms
{
    /// <summary>
    /// lower_bound：ソート済み列から「value 以上」の最初の位置を返す。
    /// </summary>
    /// <remarks>cmp 未指定時は <typeparamref name="T"/> が IComparable&lt;T&gt; を実装している必要あり。</remarks>
    public static int LowerBound<T>(IList<T> list, T value)
        => LowerBound(list, value, 0, list.Count);

    /// <summary>
    /// lower_bound（半開区間 [l, r) を探索）。
    /// </summary>
    /// <param name="l">探索開始インデックス（含む）。</param>
    /// <param name="r">探索終了インデックス（含まない）。</param>
    public static int LowerBound<T>(IList<T> list, T value, int l, int r, IComparer<T> cmp = null)
    {
        if (list is null) throw new ArgumentNullException(nameof(list));
        if (l < 0 || r < l || r > list.Count) throw new ArgumentOutOfRangeException();
        var c = cmp ?? Comparer<T>.Default;

        int left = l, right = r;
        while (left < right)
        {
            int mid = left + ((right - left) >> 1);
            if (c.Compare(list[mid], value) < 0) left = mid + 1;
            else right = mid;
        }
        return left;
    }

    /// <summary>
    /// upper_bound：ソート済み列から「value より大きい」最初の位置を返す。
    /// </summary>
    /// <remarks>cmp 未指定時は <typeparamref name="T"/> が IComparable&lt;T&gt; を実装している必要あり。</remarks>
    public static int UpperBound<T>(IList<T> list, T value)
        => UpperBound(list, value, 0, list.Count);

    /// <summary>
    /// upper_bound（半開区間 [l, r) を探索）。
    /// </summary>
    /// <param name="l">探索開始インデックス（含む）。</param>
    /// <param name="r">探索終了インデックス（含まない）。</param>
    public static int UpperBound<T>(IList<T> list, T value, int l, int r, IComparer<T> cmp = null)
    {
        if (list is null) throw new ArgumentNullException(nameof(list));
        if (l < 0 || r < l || r > list.Count) throw new ArgumentOutOfRangeException();
        var c = cmp ?? Comparer<T>.Default;

        int left = l, right = r;
        while (left < right)
        {
            int mid = left + ((right - left) >> 1);
            if (c.Compare(list[mid], value) <= 0) left = mid + 1;
            else right = mid;
        }
        return left;
    }
}
