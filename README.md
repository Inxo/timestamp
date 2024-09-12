
# Timestamp Converter

A command-line application in Go that converts Unix timestamp into a human-readable date and time. The program can also calculate the difference between the given `timestamp` and the current time, indicating whether the event occurred in the past or will occur in the future.

## Installation

1. Ensure you have Go installed (version 1.13+).
2. Clone the repository.
3. Run the following command to execute the program:

   ```bash
   make build 
   ```

## Usage

### Convert timestamp to date and time

To convert a Unix timestamp to a UTC date and time:

```bash
./build/timestamp <timestamp>
```

Example:

```bash
./build/timestamp 1726119156
```

Output:

```bash
Date and Time (UTC): 2024-09-11T09:32:36Z
```

### Calculate the difference between timestamp and the current time

To calculate the difference between the given `timestamp` and the current time, use the `-diff` flag:

```bash
./build/timestamp -diff <timestamp>
```

Example:

```bash
./build/timestamp -diff 1726119156
```

Output:

```bash
Difference: 5 days, 10 hours, 20 minutes, 15 seconds ahead
```

If the event is in the future, the program outputs `ahead`. If the event occurred in the past, it outputs `ago`.

## Arguments

- `<timestamp>` — The Unix timestamp to convert or calculate the difference for.
- `-diff` — A flag to output the time difference between the current time and the given `timestamp`.

## Example

1. Converting a timestamp to date and time:

   ```bash
   ./build/timestamp 1726119156
   ```

   Output:

   ```bash
   Date and Time (UTC): 2024-09-11T09:32:36Z
   ```

2. Calculating the time difference:

   ```bash
   ./build/timestamp -diff 1726119156
   ```

   Output:

   ```bash
   Difference: 5 days, 10 hours, 20 minutes, 15 seconds ahead
   ```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
