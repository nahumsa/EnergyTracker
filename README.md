# CLI Energy Tracker

This is a Command Line app that logs your energy level. It uses [storm](https://github.com/asdine/storm) in order to create a Bolt database and [cobra](https://github.com/spf13/cobra) for the command line interface. This CLI app is inpired by [gophercises](https://gophercises.com/).

The available commands are:
    
    - `log`: Logs your energy to a database at the present time. It takes as input an integer, I suggest to use 10 for the max value and 1 for the min value.
       Example: `EnergyTracker log 10`

    - `history`: Show all your logs in the format: year-month-day time  Energy:

## Installing the CLI App

In order to install you just need to run: `go install .` then the app will be named "EnergyTracker".

## Future Plans

For the future I plan on doing visualizations of the database.
