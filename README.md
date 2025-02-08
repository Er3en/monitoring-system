
# Log Monitoring Stack with Prometheus and Grafana

This project sets up a monitoring stack using **Docker Compose**, which includes **Prometheus**, **Grafana**, and a **Log Monitoring Service**. It is designed to provide a simple and efficient way to monitor logs and visualize metrics.

## Services

This stack consists of three main services:

1. **Log Monitor**: A custom service that processes and monitors logs. It exposes a web UI on port `8080`.
2. **Prometheus**: A monitoring system and time-series database that scrapes metrics from various sources. It exposes its web UI on port `9090`.
3. **Grafana**: A visualization tool that integrates with Prometheus to display metrics in the form of dashboards. It exposes its web UI on port `3000`.

## Docker Compose Configuration

The `docker-compose.yml` file defines the services and their configurations:

### Services

- **log-monitor**
  - **Build context**: The `log-monitor` service is built from the `./log-monitor` directory.
  - **Ports**: Exposes port `8080` for accessing the log monitoring UI.
  - **Restart policy**: Always restart unless explicitly stopped.
  - **Networks**: Connects to the `monitoring_network` for communication with Prometheus and Grafana.

- **prometheus**
  - **Image**: Uses the latest version of `prom/prometheus`.
  - **Ports**: Exposes port `9090` for the Prometheus web UI.
  - **Volumes**: Mounts a custom `prometheus.yml` configuration file to `/etc/prometheus/prometheus.yml`.
  - **Restart policy**: Always restart unless explicitly stopped.
  - **Networks**: Connects to the `monitoring_network`.

- **grafana**
  - **Image**: Uses the latest version of `grafana/grafana`.
  - **Ports**: Exposes port `3000` for the Grafana web UI.
  - **Depends on**: Depends on the Prometheus service to ensure Prometheus is running before Grafana starts.
  - **Restart policy**: Always restart unless explicitly stopped.
  - **Networks**: Connects to the `monitoring_network`.

### Networks

- **monitoring_network**: A bridge network that connects all services (`log-monitor`, `prometheus`, `grafana`) for communication.

## Getting Started

### Prerequisites

- **Docker** and **Docker Compose** must be installed on your machine. You can download them from the official Docker website:
  - [Install Docker](https://docs.docker.com/get-docker/)
  - [Install Docker Compose](https://docs.docker.com/compose/install/)

### Running the Stack

To run the monitoring stack:

1. Clone the repository or download the `docker-compose.yml` file.
2. Navigate to the project directory:
   ```bash
   cd /path/to/project
   ```
3. Build and start the containers:
   ```bash
   docker-compose up --build
   ```
4. Access the services in your browser:
   - **Log Monitor UI**: `http://localhost:8080`
   - **Prometheus UI**: `http://localhost:9090`
   - **Grafana UI**: `http://localhost:3000`  
     Default login:  
     - **Username**: `admin`  
     - **Password**: `admin`

### Stopping the Stack

To stop the containers, run:

```bash
docker-compose down
```

This will stop and remove the containers but preserve your data and configurations.

## License

This project is licensed under the MIT License.
