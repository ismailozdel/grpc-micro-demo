#include <opencv2/opencv.hpp>
#include <opencv2/highgui/highgui.hpp>
#include <opencv2/imgproc/imgproc.hpp>

int main()
{
    // Load the image
    cv::Mat img = cv::imread("../kelebek.bmp", cv::IMREAD_GRAYSCALE);

    // Check if image is loaded fine
    if(img.empty()){
        printf("Error opening image\n");
        return -1;
    }

    cv::Mat grad_x, grad_y;
    cv::Mat abs_grad_x, abs_grad_y, dst;

    // Gradient X
    cv::Sobel(img, grad_x, CV_16S, 1, 0, 3, 1, 0, cv::BORDER_DEFAULT);
    cv::convertScaleAbs(grad_x, abs_grad_x);

    // Gradient Y
    cv::Sobel(img, grad_y, CV_16S, 0, 1, 3, 1, 0, cv::BORDER_DEFAULT);
    cv::convertScaleAbs(grad_y, abs_grad_y);

    // Total Gradient (approximate)
    cv::addWeighted(abs_grad_x, 0.5, abs_grad_y, 0.5, 0, dst);

    // Save the image
    cv::imwrite("sobel.bmp", dst);

    return 0;
}