Package log implements a simple logging package.
It defines a type, Logger, with methods for formatting output.
That logger writes to standard error and prints the date and time of each logged message.
The Fatal functions call os.Exit(1) after writing the log message. The Panic functions call panic after writing the log message.
