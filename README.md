# FilterFeet

FilterFeet is a bulk image processing program that leverages all available CPU power to apply filters to a directory of photos. It operates through a command-line interface (CLI).

## Usage

To use FilterFeet, run the following command:

```
./filterfeet -src /imgs -dst /output -filter (grayscale or blur) -task (channel or waygroup) -poolsize 2
```

Replace the parameters as follows:
- `-src`: Source directory containing the images.
- `-dst`: Destination directory to store the output images.
- `-filter`: Choose the filter to apply (`grayscale` or `blur`).
- `-task`: Specify the task type (`channel` or `waygroup`).
- `-poolsize`: Set the CPU pool size (e.g., `2`).

## Example

Suppose you have a directory named `/imgs` with several images and you want to apply the `grayscale` filter using the `channel` task type with a CPU pool size of `2`. You can use the following command:

```
./filterfeet -src /imgs -dst /output -filter grayscale -task channel -poolsize 2
```


## Notes

- FilterFeet will automatically create the destination directory if it doesn't exist before processing the images.
- It utilizes all available CPU resources to expedite the image processing tasks.

Please refer to the documentation or use the `-help` flag for additional information and available options.
