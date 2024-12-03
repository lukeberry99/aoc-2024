package files

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadColumns(t *testing.T) {
	t.Run("Take a file name relative to the current file and return a slice of slices containing the columns", func(t *testing.T) {
		expected := [][]string{
			{"3", "4", "2", "1", "3", "3"},
			{"4", "3", "5", "3", "9", "3"},
		}
		actual := ReadColumns("./sample-columns.txt", 2)

		assert.Equal(t, expected, actual)
	})
}

func TestReadLines(t *testing.T) {

	t.Run("Take a file name relative to the current file and return a string array of the lines", func(t *testing.T) {

		expected := []string{"this", "is", "a", "test", "file"}
		actual := ReadLines("./sample-input.txt")

		assert.Equal(t, expected, actual)
	})

}

func TestRead(t *testing.T) {
	t.Run("Take a file name relative to the current file and return a string representation of the contents", func(t *testing.T) {
		expected := "this\nis\na\ntest\nfile\n"
		actual := Read("./sample-input.txt")

		assert.Equal(t, expected, actual)
	})
}

func TestReadLinesWithGaps(t *testing.T) {

	t.Run("Take a file name relative to the current file and return lines grouped when broken by space", func(t *testing.T) {

		expected := [][]string{
			{"this", "is", "a"},
			{"test", "file"},
			{"end"},
		}
		actual := ReadParagraphs("./input-with-gaps.txt")

		assert.Equal(t, expected, actual)
	})

}
