using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Advent18.Days
{
    class Day2
    {
        public static string file = @"D:\go\src\github.com\Qudaci\Advent\2018\Advent18\Advent18\Inputs\Day2.txt";

        public static int One()
        {
            int twos = 0;
            int threes = 0;

            try
            {
                string[] lines = System.IO.File.ReadAllLines(file);
                foreach (string line in lines)
                {
                    var charArray = line.ToCharArray();
                    bool two = false;
                    bool three = false;
                    while (charArray.Length > 0)
                    {
                        char c = charArray[0];
                        int n = charArray.Count(x => (x == c));
                        if (n == 2)
                            two = true;
                        if (n == 3)
                            three = true;

                        charArray = charArray.Where(x => (x != c)).ToArray();
                    }
                    if (two)
                        twos++;
                    if (three)
                        threes++;
                }
            }
            catch (Exception e)
            {
                System.Diagnostics.Debug.WriteLine(e.Message);
            }

            return twos*threes;
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
