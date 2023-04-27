package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/chai2010/webp"
	"github.com/google/uuid"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strings"
)

func makeWork(base64Images ...string) <-chan string {
	out := make(chan string)
	go func() {
		for _, base64Img := range base64Images {
			out <- base64Img
		}
		close(out)
	}()
	return out
}

func pipeline[I any, O any](quit <-chan struct{}, in <-chan I, f func(I) O) <-chan O {
	out := make(chan O)
	go func() {
		defer close(out)
		for v := range in {
			select {
			case out <- f(v):
			case <-quit:
				return
			}

		}
		//close(out)
	}()
	return out
}

// pipelines
// each stage has an input and an output channel
// -> stage1-> stage2-> stage 3-> ... ->

// fan-in: multiple inputs, one output
// goroutine 1 -> |
// goroutine 2 -> |
// ...         -> +  fan-in ->

// close / cancellation
// close a channel to signal that no more values will be sent on it
// goroutine 1 -> c -> |
// goroutine 2 -> c -> |
// ...         -> c -> +  fan-in ->

// request quit
// listens on a different goroutine to quit

func base64ToRawImage(base64Img string) image.Image {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64Img))
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func encodeToWebp(img image.Image) bytes.Buffer {
	var b bytes.Buffer
	if err := webp.Encode(&b, img, &webp.Options{Lossless: true}); err != nil {
		log.Fatal(err)
	}
	return b
}

func saveToDisk(imgBuf bytes.Buffer) string {
	// save to disk
	filename := fmt.Sprintf("%v.webp", uuid.New().String())
	os.WriteFile(filename, imgBuf.Bytes(), 0644)
	return filename
}

func main() {
	// Load data into the pipeline
	base64Images := makeWork("iVBORw0KGgoAAAANSUhEUgAAAOEAAADhCAMAAAAJbSJIAAAAh1BMVEX///8CAgIBAQEAAADj4+Pe3t7p6env7+/n5+f19fX7+/v09PTs7OzNzc0qKir8/PxdXV3Y2Ni2trY/Pz+goKCKioo8PDx9fX1EREQiIiKYmJhiYmLDw8NsbGw1NTUODg51dXVRUVGPj4+pqakVFRVMTEy/v78vLy+kpKR6enocHBxgYGBXV1chqfrSAAALaUlEQVR4nO1diZaiOhAdKRA3UFFUUNz39v+/74EgslQCdJMFn3fOmZ5BPCfVN6lUasu/f1988cUXn4a2JnoEjGGtzelQ9CAYYjgCH4uO6HGwwnBsAigtX8ZVT/RYmKBrOb58wR+AqS56NAzQdwMBQwDMP2+m9u0jtJSQQ0UB8FTRI6oZvVuwBGMJfRZ/PoxFD6CVAsCpL3pQdcIDiMiLfwCsRY+qRsyzDEYsfsqm0ZtuXkomwWGgbrZd0WOrBUPLjJVMSsIWbNyPYNGYxPtgbqI6lujR1QB15y9CnEOfxbUhenx/Rje3T2S0TeN3ftffJxQihx+gbVY0BkMWp6LH+CcMIK1GcxwGO3+Tl2J3WSyhAuZA9Dh/j2CrLwTArbG74uoCecryD2BzED3SX0JdFqmZEArMmrkUe2641RdyGGwZjfRqnKMTU7GEDd0y1EWpKRppG7OB83QPRMqQB+A17sRvlNMy7y2jcfp0ATTK8g/AbIsecjVMoaKE/jwVPeZKaF8qzdFwno5Fj7oKtolTb0kOFZiIHnUFnM3KFPoktpqzKfYR92ghh/5KPDXGDW45v6AwIHEveuQl0fFQ92jxA1g0RNlY8FsJoRn+0/buV3P0OU/XjSDRAoKLu/gBwL4Bnjf9NztFTKLSABKtrOeiAoc+i7b8K9H5m4QgvQ/cIoZhSs5TW7QERVjkfDOVOPR3DMl1zbmMh5ROoitaBjqWef9aNQ79o7BoGah4+YD/ImHrKloKGuxq3hkMCjxES0GBtgOElWocKnBZiZaDjANdzyjwAlXVSLxhDLaYjzTeBgA2i6W3XZ6c4N9EDhXYSZtHbJBNUl8kxza0zkDXBx117M4oRMobixpegcChL87F6icszmHXOAHp7RbMJXWAqyfURxpMUAdhZTUBgl6Ci6RRDJIjH+CB+rN7W4JiktXF37NxHykc5yRT03VwCwg8KV38hKMvbGyyLX3FvXJwlPIgbKB+boAbJb77ihNnOFTgIGNRxhzzAvtrkOrm7d9Qrw78yKhNj9iEA+dM/9YYDRUDSBjYVzFPPsA29ZI2nc+vamoGumh6LUiYlmmjXBwTc7RvmaFVaq4Sukc/oV8c8ZegCBPM0kwOVIvjNb72SewGNmajAvCXoAB6C6XifRDSRm+DAMB7c6th+74iX7zUwpmIl5w+h1TNTMK5vUbZl+4IdUMlnMWfG5mamUQCzRaVcCFCChrQcMxbk3a99OcAbkziAf3qRogYZGhrVOfHcetgmaZNz13s3DZQDmVLrD2gBuZ7V2tnXeFwj1WJiu/5kgWEbTQq+pZQy5oDieSSNs7hUowkJGxxCeNznpo1yxPHBw312cBJKuN7MMJPQfFMG+Sq8xbxOlzh35UrHvys/EE4jDO5+o8sh3Z8epjiHMpVMUTIL0nEIKw0iZAwrZeEk7NUqmaKx+4BYv9a55G2aRJ5zw7OIcyFiELAnlC8Be8oi3F5u9YAdu9FZrRw77BUyYr9OcG9m8zGWy0iL7D/45HYzrckF91DooC3uiRxmLRMxvY6PB+u7cTY1RkhvwgmElk1Y/QU+6RrmdjVequrvd3uD6mR2+ix66mmJFKmqwupiBLu6WF2B4N0OokxI+UXwVEiv/AZ9UKFJJ6oG3eHXIApVRGGRS6E9Y/zlNK03p78TUWmnAWKhIHHguii7+6Bkl8k0+kCPcO+Z9uWoPZ1etgfiPEO7ui5tEyvYC2iej92vpE4pM1vvuijvtI3fLU4zembzmFWkLgBS2m2fH1O5fBpZO5cI7lNjKfLFjEG/OLwIU08f7AtkjCIc8+W9sFQdb0ztmxvsQmtVKqEO2lOiG2PPt0ihQNwN9c+TCcQuDiFsWAr5YnOspDDF5ERlBJvt2AhjWHawc+wf4VEpndZDss8SHIoTwE0Kw6lkpAJh/LMUl+XMpFQHk0TZOyxmKXy7Bb6jQ2H8uz4RXbpbzkcSWO1seLwIQuH/emdjYTgySFiG+2oV8s0hV1BvhEXjEfpwFp9HCqKAmvxpQnWghWDIY0b0TkZ02PWCVUnh+FiFJniNpjnu7HVLKF/5BJo2oxLtkn6GxS4WIIC3qsgle0XrFR+G5SrEL+iRXTl1w4Am/9i7D/zQn/LStW3Aba8Lbj2fPPHctiqLD746pu2RyxcYsJhwOKaZ0lbe1mhMUQ9EvoiHvmJ2J1znaER/F2DmyE+pQTTmHHYCo5TnEppKjYqq5NFPik2vVO1cu36OPTXIpelOK1YkF6jhC1YcGib0T6JUDMRuGQvuK2KJfd1cqhwuFGhjWeSciPxyJxEy6lacl8vh+AxXonDPal4l4+E7DvyqmInKYca4aDkXiiHCus+S2dR9sybxB+mOSjDK6HknhuHrC+M0EnJwBw53DC13MJYqFgOgamnv/OQQEKmqcPqQ/gsbTGXUDSHjMu8Oz/COWQsYZijJ3gdMtU03b14CRm73CzxNg1j03u1EWyXMs8cxht1cOWQseXd9gRzyN5T4wr0Qz0pvLOuhjp8+hn/32otWMIf1p59wZapAjfGAj5LPgVyyKPVwl7ono9276sZ8e3aYjjkcOmlNhPr1ecQmhGpavj0A9mKjK5xaa50JRVf8+Bww6O+WzsK5PDIpSTRFCghn3bmAlUNp17fbikO0QpRytuIuYtwyCdjeFyCw0DA8sYPwN0p9TZbL9QbtCL68C+AzUqblEggjt52zv/mZe70BE43mMyKft0A60HYEaIUMbALTLHrpvBtbv0i7QIOoRVmn3dds8ydzpuoE6Y1K3qb26UCBl1CcOIuV+fgUmf6vINTXFJhPApKN7h1ytBp00kBM5GV3ZnS+ssHK9BOJBxq1Kvm/VnKK4u2P6M1+jCtlDoYB1MVz0b1n8+NVI56UGFE5pBfd+HejcLK3ciUDvTU6yWqUH+/Flau21o2w6nrUliENbcO0RaxpRAA1ty6Zywhh5PVw8ooLBLjQSsQbkUJGuFqNYAJcQzq1FtcTMeHeZn87MkryriTcsg59jFX8R0RWsuChaJr47FW1JNFO+EzleddrPoW4xDuNdV9qF62n2u4zTr8ai6Gh7xv3x+AW5eqa9tYPQ6MON6YsMprA7hb9ekBfYpE8bhejmiss79jaGV3iT+hl08v49uJXr+lOYT6b/UbQ+awAUeudUHWPanRARb1WxvtScq+oV+UUT+GSec+KEsWOmCQvIeGc93Tv7DXGkR2zN1GBWxrJeUeaB30ZKvv7y+NBgL6RHYPi8j82lmYudi+ejvvUGIHa1vb3dJFX+we1i8bbySiMn9s3U6zhWehg9O9cGD5Rm1paNfwRUKXD8OdBPx5lqDWbbqqqQQbLLwKwR/64oZS/ETvPD9tohdHuAum69t547LTnSsiRRTIeJ95B+T30Le8iRO3NQOQ8e4cGt6nj3Ad3R+2NY6E6GrW/seJ7gmMLZbGXB0fIWWS5A+H+WsQ5b8IOANAU/ze/ehyRxOQpt1OSeASUh40UELsEEtBAyX8cvjlUHZ8OfwEDj9fwu8s/XIoO/4PHH6+hO8JSMzuT37QQAljhoAQi4s+aD6HYD7uhGjZIxGla6CEEVdwnKoLNKoLLe28hgZzGA184/b+uWhUN7jh0jJfkcIGejGimfhMgUHDgc/Mimdi/PN/0lxIUhKHyNc0ff0vx2HU9Ol8Cd+cS3XdYRnsgzQa81Udkb9bDzaRD9UIogPOtmnuUh8re27HEb/kbULPHwrsX05u1X/R4pR0yA79bGUmXJqmO4tgZW48BrfxrGWQ6Tch0c0OtSFVBQ6wb5zqLIS6TPTklejqihpxeEfl+RT3cEfnJz5KwEyOhvl143CE13FKqmti64O+e8V7702zssviEKYc8qp8EYHgNBic7Jt2jCiPcXiMkOiO2NqhepPZ6BO3wi+++OKLD8d/K26zhpOeFyIAAAAASUVORK5CYII=")

	quit := make(chan struct{})
	var signal struct{}
	// decode base64 into image format
	rawImages := pipeline(quit, base64Images, base64ToRawImage)
	// encode as webp

	quit <- signal

	webpImages := pipeline(quit, rawImages, encodeToWebp)
	// save to disk
	filenames := pipeline(quit, webpImages, saveToDisk)

	for name := range filenames {
		fmt.Println(name)
	}
}
