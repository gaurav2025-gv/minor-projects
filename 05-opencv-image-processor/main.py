import cv2
import sys
import os


def process_image(image_path):

    if not os.path.exists(image_path):
        print("Error: File not found!")
        return

    image = cv2.imread(image_path)

    if image is None:
        print("Error: Could not read image!")
        return

    print("Image loaded successfully!")

    height, width = image.shape[:2]

    print(f"Original Size: {width}x{height}")

    resized = cv2.resize(image, (800, 600))

    gray = cv2.cvtColor(resized, cv2.COLOR_BGR2GRAY)

    blurred = cv2.GaussianBlur(gray, (5, 5), 0)

    edges = cv2.Canny(blurred, 100, 200)

    cv2.imwrite("output_resized.jpg", resized)
    cv2.imwrite("output_gray.jpg", gray)
    cv2.imwrite("output_blur.jpg", blurred)
    cv2.imwrite("output_edges.jpg", edges)

    print("\nProcessing complete!")
    print("Generated files:")
    print("- output_resized.jpg")
    print("- output_gray.jpg")
    print("- output_blur.jpg")
    print("- output_edges.jpg")


def main():

    if len(sys.argv) < 2:
        print("Usage: python main.py <image-file>")
        return

    image_path = sys.argv[1]

    process_image(image_path)


if __name__ == "__main__":
    main()