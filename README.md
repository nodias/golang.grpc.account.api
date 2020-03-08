# gRPC + HTTP/REST + swagger transcoding
gRPC ecosystem을 이용하여 계정관련 microservices 구축하기 예제 작성 (및 괜찮으면 생각 중인 어플 개발에 적용 할지도?)
api는 golang, front는 react로 작성 예정.

# 평소 gRPC로 빠른 통신을 할 수 있으며, 외부에 api 공유시엔 rest 및 swagger로 공유가 가능하다.


# golang.grpc.account.api
user api

# 20.03.08 진행현황 수기
make 파일을 꽤 상세히 작성하고 있습니다. make 파일은 본연의 목적보단, 지금은 세팅 가이드 목적으로 명령어를 기록하는 중입니다. (레퍼런스가 많지 않고, 거의 대부분이 설명이 좀 부족해서 시간이 좀 걸렸다.) 나중에 여러 api들이 더 만들어 질때 요긴하게 쓰일 예정이고, 블로그 또는 git에 따로 정리 하면서 정리 할 예정입니다. 아직 clean 아키텍쳐를 적용해보려고 하는 중이고 mongoDB ORM 적용 후 REST부분 refactoring 후에 web과 연결 하면서 작성 예정입니다. 
