package main

type Matrix2D [][]uint8
type Matrix3D [][][]uint8
type RuneMatrix2D [][]rune
type RuneMatrix3D [][][]rune

type ImageRequest struct {
	Base64Image       string `json:"base64Image"`
	GradientThreshold uint8  `json:"gradientThreshold"`
}

type ImageResponse struct {
	OriginalMatrix      Matrix3D `json:"originalMatrix"`
	DesaturatedMatrix   Matrix3D `json:"desaturatedMatrix"`
	DownsampledMatrix   Matrix3D `json:"downsampledMatrix"`
	GaussiansDiffMatrix Matrix3D `json:"gaussiansDiffMatrix"`
	SobelMatrix         Matrix3D `json:"sobelMatrix"`
	GradientMatrix      Matrix3D `json:"gradientMatrix"`
}

type Base64ImageResponse struct {
	OriginalImage      string `json:"originalImage"`
	DesaturatedImage   string `json:"desaturatedImage"`
	DownsampledImage   string `json:"downsampledImage"`
	GaussiansDiffImage string `json:"gaussiansDiffImage"`
	SobelImage         string `json:"sobelImage"`
	GradientImage      string `json:"gradientImage"`
}
