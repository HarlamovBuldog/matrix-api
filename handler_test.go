package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func requestFactory(t *testing.T, url, testFileName string) *http.Request {
	testFilePath := path.Join("test_data", testFileName)
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	testFile, err := os.Open(testFilePath)
	require.NoError(t, err)
	t.Cleanup(func() {
		testFile.Close()
	})
	part, err := writer.CreateFormFile("file", filepath.Base(testFilePath))
	require.NoError(t, err)
	_, err = io.Copy(part, testFile)
	require.NoError(t, err)
	err = writer.Close()
	require.NoError(t, err)
	req := httptest.NewRequest("GET", url, buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req
}

func TestHandler_Echo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := requestFactory(t, "/echo", "matrix.csv")
		echo(w, req)
		actualCode := w.Result().StatusCode
		require.Equal(t, http.StatusOK, actualCode)
		expectedResp := `1,2,3
4,5,6
7,8,9` + "\n"
		actualResp := &bytes.Buffer{}
		_, err := actualResp.ReadFrom(w.Result().Body)
		require.NoError(t, err)
		w.Result().Body.Close()
		require.Equal(t, expectedResp, actualResp.String())
	})
	t.Run("fail: invalid input", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := requestFactory(t, "/echo", "not_all_int.csv")
		echo(w, req)
		actualCode := w.Result().StatusCode
		require.Equal(t, http.StatusBadRequest, actualCode)
		expectedResp := "bad request: all matrix elems should be numbers"
		actualResp := &bytes.Buffer{}
		_, err := actualResp.ReadFrom(w.Result().Body)
		require.NoError(t, err)
		w.Result().Body.Close()
		require.Equal(t, expectedResp, actualResp.String())
	})
}

func TestHandler_Invert(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := requestFactory(t, "/invert", "matrix.csv")
		invert(w, req)
		actualCode := w.Result().StatusCode
		require.Equal(t, http.StatusOK, actualCode)
		expectedResp := `1,4,7
2,5,8
3,6,9` + "\n"
		actualResp := &bytes.Buffer{}
		_, err := actualResp.ReadFrom(w.Result().Body)
		require.NoError(t, err)
		w.Result().Body.Close()
		require.Equal(t, expectedResp, actualResp.String())
	})
	t.Run("fail: invalid input", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := requestFactory(t, "/invert", "not_all_int.csv")
		invert(w, req)
		actualCode := w.Result().StatusCode
		require.Equal(t, http.StatusBadRequest, actualCode)
		expectedResp := "bad request: all matrix elems should be numbers"
		actualResp := &bytes.Buffer{}
		_, err := actualResp.ReadFrom(w.Result().Body)
		require.NoError(t, err)
		w.Result().Body.Close()
		require.Equal(t, expectedResp, actualResp.String())
	})
}

func TestHandler_Flatten(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := requestFactory(t, "/flatten", "matrix.csv")
		flatten(w, req)
		actualCode := w.Result().StatusCode
		require.Equal(t, http.StatusOK, actualCode)
		expectedResp := `1,2,3,4,5,6,7,8,9` + "\n"
		actualResp := &bytes.Buffer{}
		_, err := actualResp.ReadFrom(w.Result().Body)
		require.NoError(t, err)
		w.Result().Body.Close()
		require.Equal(t, expectedResp, actualResp.String())
	})
	t.Run("fail: invalid input", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := requestFactory(t, "/flatten", "not_all_int.csv")
		flatten(w, req)
		actualCode := w.Result().StatusCode
		require.Equal(t, http.StatusBadRequest, actualCode)
		expectedResp := "bad request: all matrix elems should be numbers"
		actualResp := &bytes.Buffer{}
		_, err := actualResp.ReadFrom(w.Result().Body)
		require.NoError(t, err)
		w.Result().Body.Close()
		require.Equal(t, expectedResp, actualResp.String())
	})
}

func TestHandler_Sum(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := requestFactory(t, "/sum", "matrix.csv")
		sum(w, req)
		actualCode := w.Result().StatusCode
		require.Equal(t, http.StatusOK, actualCode)
		expectedResp := "45\n"
		actualResp := &bytes.Buffer{}
		_, err := actualResp.ReadFrom(w.Result().Body)
		require.NoError(t, err)
		w.Result().Body.Close()
		require.Equal(t, expectedResp, actualResp.String())
	})
	t.Run("fail: invalid input", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := requestFactory(t, "/sum", "not_all_int.csv")
		sum(w, req)
		actualCode := w.Result().StatusCode
		require.Equal(t, http.StatusBadRequest, actualCode)
		expectedResp := "bad request: all matrix elems should be numbers"
		actualResp := &bytes.Buffer{}
		_, err := actualResp.ReadFrom(w.Result().Body)
		require.NoError(t, err)
		w.Result().Body.Close()
		require.Equal(t, expectedResp, actualResp.String())
	})
}

func TestHandler_Multiply(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := requestFactory(t, "/multiply", "matrix.csv")
		multiply(w, req)
		actualCode := w.Result().StatusCode
		require.Equal(t, http.StatusOK, actualCode)
		expectedResp := "362880\n"
		actualResp := &bytes.Buffer{}
		_, err := actualResp.ReadFrom(w.Result().Body)
		require.NoError(t, err)
		w.Result().Body.Close()
		require.Equal(t, expectedResp, actualResp.String())
	})
	t.Run("fail: invalid input", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := requestFactory(t, "/multiply", "not_all_int.csv")
		multiply(w, req)
		actualCode := w.Result().StatusCode
		require.Equal(t, http.StatusBadRequest, actualCode)
		expectedResp := "bad request: all matrix elems should be numbers"
		actualResp := &bytes.Buffer{}
		_, err := actualResp.ReadFrom(w.Result().Body)
		require.NoError(t, err)
		w.Result().Body.Close()
		require.Equal(t, expectedResp, actualResp.String())
	})
}
