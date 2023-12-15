# FilterFleet

FilterFleet is a bulk image processing program that leverages all available CPU power to apply filters to a directory of photos. It operates through a command-line interface (CLI).

![This is an image](https://github.com/gildasgatel/FilterFeet/blob/master/asset/filterfleet.jpg)


## Usage

To use FilterFleet, run the following command:

```
./FilterFleet -src /imgs -dst /output -filter (grayscale or blur) -task (channel or waitgroup) -poolsize 2
```

Replace the parameters as follows:
- `-src`: Source directory containing the images.
- `-dst`: Destination directory to store the output images.
- `-filter`: Choose the filter to apply (`grayscale` or `blur`).
- `-task`: Specify the task type (`channel` or `waitgroup`).
- `-poolsize`: Set the CPU pool size (e.g., `2`).

## Example

Suppose you have a directory named `/imgs` with several images and you want to apply the `grayscale` filter using the `channel` task type with a CPU pool size of `2`. You can use the following command:

```
./FilterFleet -src /imgs -dst /output -filter grayscale -task channel -poolsize 2
```
![This is an image](https://github.com/gildasgatel/FilterFeet/blob/master/asset/process.png)


## Notes

- FilterFleet will automatically create the destination directory if it doesn't exist before processing the images.
- It utilizes all available CPU resources to expedite the image processing tasks.

