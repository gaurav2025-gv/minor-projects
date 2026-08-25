import cv2
import os
video_path="input.mp4"
output_folder="frames"

os.makedirs(output_folder,exist_ok=True)

cap=cv2.VideoCapture(video_path)
if not cap.isOpened():
    print("Error: Could not open video")
    exit()

fps = int(cap.get(cv2.CAP_PROP_FPS))
frame_count = int(cap.get(cv2.CAP_PROP_FRAME_COUNT))
width = int(cap.get(cv2.CAP_PROP_FRAME_WIDTH))
height = int(cap.get(cv2.CAP_PROP_FRAME_HEIGHT))

print("FPS:", fps)
print("Total Frames:", frame_count)
print("Resolution:", width, "x", height)

sample_rate = 30
frame_number = 0
saved_count = 0

while True:
    ret, frame = cap.read()

    if not ret:
        break

    if frame_number % sample_rate == 0:
        filename = f"{output_folder}/frame_{saved_count + 1:03d}.jpg"
        cv2.imwrite(filename, frame)

        saved_count += 1
        print("Saved:", filename)

    frame_number += 1

cap.release()

print("\nDone!")
print("Frames saved:", saved_count)