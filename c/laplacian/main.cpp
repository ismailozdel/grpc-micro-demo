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

    cv::Mat laplacian;

    // Calculate Laplacian
    cv::Laplacian(img, laplacian, CV_16S, 3, 1, 0, cv::BORDER_DEFAULT);
    cv::convertScaleAbs(laplacian, laplacian);

    // Save the image
    cv::imwrite("laplacian.bmp", laplacian);

    return 0;
}