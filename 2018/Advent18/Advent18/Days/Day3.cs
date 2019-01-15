using System;
using System.Text.RegularExpressions;
using System.IO;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Advent18.Days
{
    class Square
    {
        int left;
        int top;
        int width;
        int height;

        private static Regex rx = new Regex(@"#\d+ @ (\d+),(\d+): (\d+)x(\d+)");

        public Square(string input)
        {
            MatchCollection matches = rx.Matches(input);
            if (matches.Count != 1 || matches[0].Groups.Count != 5)
                throw new Exception("incorrect input");


            this.left = Int32.Parse(matches[0].Groups[1].Value);
            this.top = Int32.Parse(matches[0].Groups[2].Value);
            this.width = Int32.Parse(matches[0].Groups[3].Value);
            this.height = Int32.Parse(matches[0].Groups[4].Value);
        }
    }

    class Day3
    {
        public static string file = @"D:\go\src\github.com\Qudaci\Advent\2018\Advent18\Advent18\Inputs\Day3.txt";

        public static int One()
        {
            List<Square> plans = new List<Square>();
            List<Square> overlap = new List<Square>();
            List<Square> overoverlap = new List<Square>();

            try
            {
                string[] lines = System.IO.File.ReadAllLines(file);
                foreach (string line in lines)
                {
                    Square sq = new Square(line);
                    plans.Add(sq);

                }
            }
            catch (Exception e)
            {
                System.Diagnostics.Debug.WriteLine(e.Message);
            }

            return 0;
        }

        public static string Two()
        {
            int twos = 0;
            int threes = 0;

            try
            {
                string[] lines = System.IO.File.ReadAllLines(file);
                for (int i = 0; i < lines.Length - 1; i++)
                {
                    var lineA = lines[i].ToCharArray();
                    for (int j = i + 1; j < lines.Length; j++)
                    {
                        var lineB = lines[j].ToCharArray();
                        bool dif = false;
                        int difPos = 0;
                        for (int pos = 0; pos < lineA.Length; pos++)
                        {
                            if (lineA[pos] == lineB[pos])
                                continue;

                            if (!dif)
                            {
                                dif = true;
                                difPos = pos;
                            }
                            else
                            {
                                dif = false;
                                break;
                            }
                        }
                        if (dif)
                        {
                            var val = lineA.Take(difPos).Concat(lineB.Skip(difPos+1)).ToArray();
                            return new string(val);
                        }
                    }
                }
            }
            catch (Exception e)
            {
                System.Diagnostics.Debug.WriteLine(e);
            }
            return "nothing found";
        }
    }
}
