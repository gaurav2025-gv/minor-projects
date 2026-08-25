from ultralytics import YOLO
import cv2

# Load YOLOv8 model
model = YOLO("yolov8n.pt")

# Read input image
image = cv2.imread("road.jpg")

# Run object detection
results = model(image)

# Process each detection result
for result in results:
    boxes = result.boxes

    for box in boxes:

        # Get bounding box coordinates
        x1, y1, x2, y2 = map(int, box.xyxy[0])

        # Get confidence score
        confidence = float(box.conf[0])

        # Get class ID
        class_id = int(box.cls[0])

        # Convert class ID to class name
        class_name = model.names[class_id]

        # Draw bounding box
        cv2.rectangle(
            image,
            (x1, y1),
            (x2, y2),
            (0, 255, 0),
            2
        )

        # Draw label and confidence
        cv2.putText(
            image,
            f"{class_name} {confidence:.2f}",
            (x1, y1 - 10),
            cv2.FONT_HERSHEY_SIMPLEX,
            0.6,
            (0, 255, 0),
            2
        )

# Save processed image
cv2.imwrite("output.jpg", image)

print("Detection complete!")