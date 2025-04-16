# 🧠 Mood Tracker API

A simple emotion-tracking RESTful API built with Go, PostgreSQL, and Render for cloud deployment.

## 🚀 Features

- ✅ Add new emotions
- 📖 Get all emotions
- ❌ Delete emotions
- 🛠️ Easily extendable (update, filter, etc.)
- 🌐 Ready for cloud deployment (Render / Railway)

## 🧪 Tech Stack

- Go (Golang)
- Fiber Web Framework
- GORM ORM
- PostgreSQL
- Render (for deployment)
- docker

## 🌐 Deployment

This project is deployed on **Render**. You can access the live API [here](https://mood-tracker-jfus.onrender.com). Might shut down when you visit because it's a free plan.

## 📦 Installation

1. Clone the repository:

   ```bash
   git clone path
   cd mood-tracker
   ```

2. **Set up environment variables**:

   If you're running locally, create a `.env` file in the root directory with the following content:

   ```bash
   DB_URL= Link to your db
   ```

   For Render deployment, you can set the environment variables directly in the **Render Dashboard** by adding the following environment variables:

   - **DB_URL**: Your PostgreSQL database connection link

3. Install Go dependencies:

   ```bash
   go mod tidy
   ```

4. **Run the app locally**:

   If you are running the app locally, simply use the following command:

   ```bash
   go run main.go
   ```

   The app will be accessible at `http://localhost:8080`.

## 🌍 Cloud Deployment (Render)

If you'd like to deploy this API to **Render**:

1. Create a Render account at [Render.com](https://render.com).
2. Create a new **Web Service** from the Render Dashboard.
3. Link your GitHub repository for automatic deployments.
4. Configure the **Environment Variables** in the Render Dashboard to match the ones used in the `.env` file (DB_URL).
5. Once deployed, you can access your API at the URL provided by Render.

## 🧰 API Endpoints

### `POST /emotions`

- **Description**: Adds a new emotion to the database.
- **Body**: 
  ```json
  {
    "emotion": "happy",
    "description": "Feeling great!"
  }
  ```

- **Response**:
  ```json
  {
    "id": 1,
    "emotion": "happy",
    "description": "Feeling great!",
    "created_at": "2023-12-01T00:00:00Z",
    "updated_at": "2023-12-01T00:00:00Z"
  }
  ```

### `GET /emotions`

- **Description**: Fetches all stored emotions.
- **Response**:
  ```json
  [
    {
      "id": 1,
      "emotion": "happy",
      "description": "Feeling great!",
      "created_at": "2023-12-01T00:00:00Z",
      "updated_at": "2023-12-01T00:00:00Z"
    }
  ]
  ```

### `DELETE /emotions/:id`

- **Description**: Deletes an emotion by its ID.
- **Response**:
  ```json
  {
    "message": "Emotion deleted successfully"
  }
  ```

## 🛠️ Development

### Local Development with Docker

To run the project using Docker, follow these steps:

1. **Build the Docker image**:

   ```bash
   docker build -t mood-tracker .
   ```

3. **Run the Docker container**:

   ```bash
   docker run -p 8080:8080 mood-tracker
   ```

Your application will now be accessible at `http://localhost:8080`.

## 👨‍💻 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```