Certainly! Here's a README.md file for your project:

# Arithmetic Progression Queue Service

The Arithmetic Progression Queue Service is a tool for managing and calculating arithmetic progressions in a queue. Tasks are added to the queue and processed in the order they are received. The service allows for parallel execution of tasks, enhancing performance and efficiency.

Also, it is written following the go style guide, but in exactly the amount required for this project to prevent overengineering 

Below are key features and instructions for using the service.

## Key Features

1. **Queue-Based Processing**: Tasks are added to a queue and executed sequentially, ensuring organized and orderly task execution.

2. **Parallel Task Execution**: The service supports parallel task execution to improve throughput and efficiency. Users can configure the maximum number of parallel tasks (N) according to their requirements.

3. **Task Management**: Each task is associated with essential parameters, including the initial value, common difference, and other relevant details. The service accurately tracks task progress.

4. **Status Tracking**: Users can monitor task statuses, such as whether they are in the queue, in progress, or completed. Task details are readily accessible for transparency.

5. **Time-to-Live (TTL) Management**: The service handles TTL for tasks, automatically removing completed tasks from the queue after a specified time to free up resources.

6. **Error Handling**: The service includes error-handling mechanisms to validate and process task parameters, providing feedback to users when input data is incorrect.

## Usage

1. **Build the Service**: Use the `make build` command to build the service binary.
   ```bash
   make build
   ```

2. **Run the Service**: Execute the service with the desired number of parallel tasks (N).
   ```bash
   ./2queue-service --maxParallelTasks N
   ```

3. **Add a Task**: To add a task to the queue, make a POST request to `/addTask` with the required parameters. (**Or in Postman ( in Body `form-encode`**)). 
   ```bash
   curl -X POST "http://localhost:8085/addTask?n=8&d=2.0&n1=5.0&I=2.5&TTL=5"
   ```

4. **Get Task Status**: Retrieve the status of all tasks in the queue by making a GET request to `/getTasks`.
   ```bash
   curl "http://localhost:8085/getTasks"
   ```

5. **Error Handling**: The service validates task parameters and returns appropriate error messages if input data is incorrect.

## Configuration

- **Parallel Tasks (N)**: Users can specify the maximum number of parallel tasks by providing the `--maxParallelTasks` flag when running the service. Adjust this value to match the system's capacity and performance requirements.

## Cleaning Up

To remove the service binary, use the `make clean` command.

```bash
make clean
```

## Authors

- [Dmitrii Kumancev](https://github.com/DmitriiKumancev)


