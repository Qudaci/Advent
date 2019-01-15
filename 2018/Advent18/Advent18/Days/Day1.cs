using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Advent18.Days
{
    class Day1
    {
        public static string file = @"D:\go\src\github.com\Qudaci\Advent\2018\Advent18\Advent18\Inputs\Day1.txt";

        public static int One()
        {
            int val = 0;
            try
            {
                string[] lines = System.IO.File.ReadAllLines(file);
                foreach (string line in lines)
                {
                    val += Int32.Parse(line);
                }
            }
            catch (Exception e)
            {
                System.Diagnostics.Debug.WriteLine(e.Message);
            }

            return val;
        }

        public static int Two()
        {
            int f = 0;
            List<int> frequencies = new List<int>() { 0 };

            try
            {
                string[] lines = System.IO.File.ReadAllLines(file);
                while (true)
                {
                    foreach (string line in lines)
                    {
                        int val = Int32.Parse(line);
                        f += val;
                        int i = frequencies.FindIndex(x => (x >= f));
                        if (i == -1)
                        {
                            frequencies.Add(f);
                        }
                        else
                        {
                            if (frequencies[i] == f)
                                return f;

                            frequencies.Insert(i, f);
                        }
                    }
                }
            }
            catch (Exception e)
            {
                System.Diagnostics.Debug.WriteLine(e.Message);
            }
            throw new System.Exception();
        }
    }
}
