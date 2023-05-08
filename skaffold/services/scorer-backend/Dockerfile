FROM python:3.9:alpine

WORKDIR /app

COPY . .

RUN pip install -r requirements.txt

EXPOSE 5000


# Run app.py when the container launches
CMD ["python", "main.py"]
