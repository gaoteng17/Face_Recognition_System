# -*- codeing: utf-8 -*-

import cv2
import sys
import os
import numpy as np
import skimage
import caffe
import sklearn.metrics.pairwise as pw

haar_dir = 'haarcascade_frontalface_default.xml'
caffe_dir = '/root/caffe/'
vgg_dir = '/root/caffe/examples/vgg_face_caffe/'
net = caffe.Classifier('deploy.prototxt', vgg_dir + 'VGG_FACE.caffemodel',caffe.TEST)
sys.path.insert(0, caffe_dir + 'python')

#GPUmode
#caffe.set_mode_gpu()





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



def getFaceFromCam(frame_num, Output):

	if Output[-1] != '/' :
		Output = Output + '/'

	if not os.path.exists(Output):
		os.makedirs(Output)

# transform the frames captured by camera to face picture(224*224)

	haar = cv2.CascadeClassifier(haar_dir)

	cam = cv2.VideoCapture(0)

	num = 1

	while True:

		_, img = cam.read()  

		gray_img = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)

		faces = haar.detectMultiScale(gray_img, 1.3, 5)

		for f_x, f_y, f_w, f_h in faces:

			#print(f_x, f_y, f_w, f_h)

			face = img[f_y:f_y+f_h, f_x:f_x+f_w]

			face = cv2.resize(face, (224,224))

			cv2.imwrite(Output + str(num) + '.jpg', face)

			num = num + 1

			cv2.rectangle(img, (f_x,f_y),(f_x + f_w,f_y + f_h), (255,0,0),3)

			cv2.imshow('image',img)

		if num > frame_num:
			return True

		key = cv2.waitKey(30) & 0xff

		if key == 27:
			return False



def comparePic(Pic1, Pic2, threshold = 0.88):
	
	avg = np.array([129.1863,104.7624,93.5940])

	img = caffe.io.load_image(Pic1)
	img = caffe.io.resize_image(img, (224,224))
	img = img[:,:,::-1]*255.0 # convert RGB->BGR
	img = img - avg # subtract mean (numpy takes care of dimensions :)
	img = img.transpose((2,0,1)) 
	img = img[None,:] # add singleton dimension

	out = net.forward(data = img)
	# print out.keys()

	feature1 = np.float64(out["fc7"])

	feature1=np.reshape(feature1,(1,4096))

	#print(feature1)

	#X=read_image(Pic2)

	img = caffe.io.load_image(Pic2)
	img = caffe.io.resize_image(img, (224,224))
	img = img[:,:,::-1]*255.0 # convert RGB->BGR
	img = img - avg # subtract mean (numpy takes care of dimensions :)
	img = img.transpose((2,0,1)) 
	img = img[None,:] # add singleton dimension

	out = net.forward(data=img)

	feature2 = np.float64(out["fc7"])
	feature2=np.reshape(feature2,(1,4096))

	predicts=pw.cosine_similarity(feature1, feature2)

	print("the similarity of %s and %s is: %f\n\n"%(Pic1,Pic2,predicts))


	if predicts >= threshold:
		return True
	else:
		return False


def test(num):
	
	pattern = './out/1.jpg'

	my1 = './noise/'

	my2 = './out/'

	other = './other_faces/'

	sum1 = 0

	sum2 = 0

	sum3 = 0

	for i in range(num):
	
		tmp = my1 + str(i+1) + '.jpg'

		if comparePic(pattern, tmp):
			sum1 = sum1 + 1

		tmp = other + '(' + str(i+1) + ').jpg'

		if not(comparePic(pattern, tmp)):
			sum2 = sum2 + 1

		tmp = my2 + str(i+1) + '.jpg'

		if comparePic(pattern, tmp):
			sum3 = sum3 + 1


	print('accuracy:\n\n')

	print("my:  %d / %d  (%f%%) \n"%(sum3, num, sum3*100.0/num))

	print("my(input noise):  %d / %d  (%f%%) \n"%(sum1, num, sum1*100.0/num))

	print("other:  %d / %d  (%f%%) \n\n\n"%(sum2, num, sum2*100.0/num))



if __name__ == '__main__':

	input_dir = sys.argv[1]
	pattern_dir = sys.argv[2]

#	getFaceFromPic('timg.jpeg', 'out.jpg')

#	getFaceFromCam(10, "./out"))

	print(comparePic("me.jpg", "noise.jpg")) #same

	print(comparePic("me.jpg", "other.jpg")) #difference

	#test(100)

	#comparePic("./out/1.jpg", "./1/1.jpg")