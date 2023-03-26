## Focuz - Your Personal Work-Break Reminder

Focuz is a command-line tool designed to help developers take regular breaks while working for prolonged periods of time. It reminds you to take a break after a certain duration of work and helps you maintain a healthy work-life balance.
Features

- Set custom work duration and break duration in minutes or seconds
- Option to interrupt work or mark work as done
- Pop-up notifications to alert you when work or break time is up
- Calculates average work and break times for analytics

## Installation

To install Focuz, you must have Go installed on your system. You can download and install Go from the official Go website.

Once you have Go installed, you can install Focuz by running the following command:

`go get github.com/raufsmestone/focuz`


## Usage

To use Focuz, navigate to the directory where it is installed and run the following command:


You will be prompted to enter your desired work duration and break duration. Once you have entered these values, Focuz will begin reminding you to take breaks after your set work duration has elapsed.

While Focuz is running, you have the option to interrupt your work or mark your work as done by typing interrupt or workdone, respectively.

## Options
Press 'i' to interrupt the current work/break session
Press 'd' to mark the current work session as done and exit
Press 'q' to quit the program

To exit Focuz, simply press Ctrl + C 


Contributing

If you'd like to contribute to Focuz, feel free to submit a pull request. Before submitting a pull request, be sure to run the tests by running the following command in the root directory:


License

Focuz is released under the MIT License. Feel free to use, modify, and distribute this tool as you see fit.
