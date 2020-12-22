package service

import "github.com/qbbang/goGin-framework/01-getting-started/entity"

// 인터페이스 정의
type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entitiy.Video
}

func New() VideoService {
	return &videoService{}
}

// 메소드 - 함수명 앞에 부가적인 파라미터를 추가한 형태
// 메소드는 데이터 고조의 특성과 동작을 표현, 사용자는 객체의 구현에 직접 접근할 필요 없이 호출만 하면 됨.
func (service *videoService) Save(video entitiy.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
