# -*- codeing: utf-8 -*-

import sys
import cv2

haar_dir = 'haarcascade_frontalface_default.xml'

def getFaceFromPic(Input, Output):

# transform the picture to face picture(224*224)

	haar = cv2.CascadeClassifier(haar_dir)

	img = cv2.imread(Input)

	gray_img = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)

	faces = haar.detectMultiScale(gray_img, 1.3, 5)

	for f_x, f_y, f_w, f_h in faces:

		#print(f_x, f_y, f_w, f_h)

		face = img[f_y:f_y+f_h, f_x:f_x+f_w]

		face = cv2.resize(face, (224,224))

		cv2.imwrite(Output, face)




if __name__ == '__main__':

	#input_dir = sys.argv[1]
	#output_dir = sys.argv[2]

	input_dir = "test.jpg"
	output_dir = "out.jpg"

	getFaceFromPic(input_dir, output_dir)
