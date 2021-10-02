# LogAnalyzer

**This example is adopt from [The Art of Unit Testing: with examples in C# 2nd Edition](https://www.amazon.com/Art-Unit-Testing-examples/dp/1617290890)**

Your company has many internal products it uses to monitor its applications at customer sites. All these products write log files and place them in a special directory. The log files are written in a proprietary format that your company has come up with that can’t be parsed by any existing third-party tools. You’re tasked with building a product, LogAn, that can analyze these log files and find special cases and events in them. When it finds these cases and events, it should alert the appropriate parties.