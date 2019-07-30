package util

import (
	"errors"
	"github.com/golang/protobuf/proto"
	sitewhere "sdkeji/go_mqtt/proto"
)

func sitewhereMarshal(deviceEventHeader sitewhere.DeviceEvent_Header, deviceEventRequest proto.Message) (protoData []byte, err error) {
	deviceHeaderData, err := proto.Marshal(&deviceEventHeader)
	if err != nil {
		return
	}
	protoData = append(protoData, byte(len(deviceHeaderData)))
	protoData = append(protoData, deviceHeaderData...)

	var deviceEventData []byte

	switch deviceEventRequest.(type) {
	case *sitewhere.DeviceEvent_DeviceLocation:
		deviceEventRequest := deviceEventRequest.(*sitewhere.DeviceEvent_DeviceLocation)
		deviceEventData, err = proto.Marshal(deviceEventRequest)
		if err != nil {
			return
		}

	case *sitewhere.DeviceEvent_DeviceAlert:
		deviceEventRequest := deviceEventRequest.(*sitewhere.DeviceEvent_DeviceAlert)
		deviceEventData, err = proto.Marshal(deviceEventRequest)
		if err != nil {
			return
		}

	case *sitewhere.DeviceEvent_DeviceAcknowledge:
		deviceEventRequest := deviceEventRequest.(*sitewhere.DeviceEvent_DeviceAcknowledge)
		deviceEventData, err = proto.Marshal(deviceEventRequest)
		if err != nil {
			return
		}

	case *sitewhere.DeviceEvent_DeviceRegistrationRequest:
		deviceEventRequest := deviceEventRequest.(*sitewhere.DeviceEvent_DeviceRegistrationRequest)
		deviceEventData, err = proto.Marshal(deviceEventRequest)
		if err != nil {
			return
		}

	case *sitewhere.DeviceEvent_DeviceMeasurement:
		deviceEventRequest := deviceEventRequest.(*sitewhere.DeviceEvent_DeviceMeasurement)
		deviceEventData, err = proto.Marshal(deviceEventRequest)
		if err != nil {
			return
		}

	case *sitewhere.DeviceEvent_DeviceStream:
		deviceEventRequest := deviceEventRequest.(*sitewhere.DeviceEvent_DeviceStream)
		deviceEventData, err = proto.Marshal(deviceEventRequest)
		if err != nil {
			return
		}

	case *sitewhere.DeviceEvent_DeviceStreamData:
		deviceEventRequest := deviceEventRequest.(*sitewhere.DeviceEvent_DeviceStreamData)
		deviceEventData, err = proto.Marshal(deviceEventRequest)
		if err != nil {
			return
		}

	default:
		err = errors.New("type assert error")
		return
	}

	protoData = append(protoData, byte(len(deviceEventData)))
	protoData = append(protoData, deviceEventData...)

	return
}

func sitwhereUnmarshal(srcData []byte, deviceEventHeader *sitewhere.DeviceEvent_Header, deviceEventRequest proto.Message) (err error) {
	if len(srcData) == 0 {
		err = errors.New("byte data is empty")
		return
	}

	headerSize := srcData[0]
	if len(srcData) < int(headerSize+2) {
		err = errors.New("length of byte data is not enough")
		return
	}

	headerData := srcData[1 : headerSize+1]
	bodyData := srcData[headerSize+2:]

	err = proto.Unmarshal(headerData, deviceEventHeader)
	if err != nil {
		return
	}

	err = proto.Unmarshal(bodyData, deviceEventRequest)
	if err != nil {
		return
	}

	return nil
}
