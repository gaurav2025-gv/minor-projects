import sqlite3

conn = sqlite3.connect("detections.db")
cursor = conn.cursor()

class_name = "pothole"
confidence = 0.87
severity = "high"
latitude = 25.4358
longitude = 81.8463

cursor.execute("""
INSERT INTO detections
(class_name, confidence, severity, latitude, longitude)
VALUES (?, ?, ?, ?, ?)
""", (
    class_name,
    confidence,
    severity,
    latitude,
    longitude
))

conn.commit()

conn.close()

print("Detection saved!")