import os
import sys
from PIL import Image


def analyze_image(filename):
    if not os.path.exists(filename):
        print("Error: File not found!")
        return

    try:
        image = Image.open(filename)

        file_size = os.path.getsize(filename)

        print("\n===== IMAGE METADATA =====")
        print(f"File Name : {os.path.basename(filename)}")
        print(f"Format    : {image.format}")
        print(f"Width     : {image.width}px")
        print(f"Height    : {image.height}px")
        print(f"Mode      : {image.mode}")
        print(f"File Size : {file_size / 1024:.2f} KB")

    except Exception as err:
        print("Error:", err)


def main():
    if len(sys.argv) < 2:
        print("Usage: python main.py <image-file>")
        return

    filename = sys.argv[1]

    analyze_image(filename)


if __name__ == "__main__":
    main()