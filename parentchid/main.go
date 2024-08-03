package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var dataStr string = `
[{"id":"cbc25911-6f3a-45fe-bac2-44a0408b6343","text":"23","file_url":null,"file_name":null,"file_description":null,"file_size":null,"file_type":null,"file_mime_type":null,"parent_uuid":"e0df428b-ee87-44c0-931f-26c5a4acff8e","created_at":"2024-05-13T03:19:51.727565Z","updated_at":"2024-05-13T03:19:51.727565Z","created_by_name":"user 50","created_by_id":"e6772dd5-9c96-46d5-a00d-eaba7019b332","created_by_photo_url":null,"childrens":null},{"id":"8e3177b9-54e7-4b46-8302-3539a65b66da","text":"22","file_url":null,"file_name":null,"file_description":null,"file_size":null,"file_type":null,"file_mime_type":null,"parent_uuid":"e0df428b-ee87-44c0-931f-26c5a4acff8e","created_at":"2024-05-13T03:19:48.994916Z","updated_at":"2024-05-13T03:19:48.994916Z","created_by_name":"user 50","created_by_id":"e6772dd5-9c96-46d5-a00d-eaba7019b332","created_by_photo_url":null,"childrens":null},{"id":"3f70dde5-9979-4970-991a-d03defd8d2d2","text":"21","file_url":null,"file_name":null,"file_description":null,"file_size":null,"file_type":null,"file_mime_type":null,"parent_uuid":"e0df428b-ee87-44c0-931f-26c5a4acff8e","created_at":"2024-05-13T03:19:46.400942Z","updated_at":"2024-05-13T03:19:46.400942Z","created_by_name":"user 50","created_by_id":"e6772dd5-9c96-46d5-a00d-eaba7019b332","created_by_photo_url":null,"childrens":null},{"id":"e0df428b-ee87-44c0-931f-26c5a4acff8e","text":"\u003cp\u003e2\u003c/p\u003e","file_url":null,"file_name":null,"file_description":null,"file_size":null,"file_type":null,"file_mime_type":null,"parent_uuid":"","created_at":"2024-05-13T03:19:34.944171Z","updated_at":"2024-05-13T03:19:34.944171Z","created_by_name":"user 50","created_by_id":"e6772dd5-9c96-46d5-a00d-eaba7019b332","created_by_photo_url":null,"childrens":null},{"id":"dd0fb817-d71d-466b-9b47-c1447ae9a8dd","text":"17","file_url":null,"file_name":null,"file_description":null,"file_size":null,"file_type":null,"file_mime_type":null,"parent_uuid":"89830116-f497-45fd-bbc3-fdda5ee7d6b7","created_at":"2024-05-13T03:19:22.878208Z","updated_at":"2024-05-13T03:19:22.878208Z","created_by_name":"user 50","created_by_id":"e6772dd5-9c96-46d5-a00d-eaba7019b332","created_by_photo_url":null,"childrens":null},{"id":"d06764ba-a543-4c1b-9ae3-18967228f255","text":"16","file_url":null,"file_name":null,"file_description":null,"file_size":null,"file_type":null,"file_mime_type":null,"parent_uuid":"89830116-f497-45fd-bbc3-fdda5ee7d6b7","created_at":"2024-05-13T03:19:20.007889Z","updated_at":"2024-05-13T03:19:20.007889Z","created_by_name":"user 50","created_by_id":"e6772dd5-9c96-46d5-a00d-eaba7019b332","created_by_photo_url":null,"childrens":null},{"id":"19f50348-671e-4e17-b48d-5f4366d0980d","text":"15","file_url":null,"file_name":null,"file_description":null,"file_size":null,"file_type":null,"file_mime_type":null,"parent_uuid":"89830116-f497-45fd-bbc3-fdda5ee7d6b7","created_at":"2024-05-13T03:19:16.492073Z","updated_at":"2024-05-13T03:19:16.492073Z","created_by_name":"user 50","created_by_id":"e6772dd5-9c96-46d5-a00d-eaba7019b332","created_by_photo_url":null,"childrens":null},{"id":"69faae71-13b3-4d7e-a633-3df8b0e63e5b","text":"14","file_url":null,"file_name":null,"file_description":null,"file_size":null,"file_type":null,"file_mime_type":null,"parent_uuid":"89830116-f497-45fd-bbc3-fdda5ee7d6b7","created_at":"2024-05-13T03:19:03.901677Z","updated_at":"2024-05-13T03:19:03.901677Z","created_by_name":"user 50","created_by_id":"e6772dd5-9c96-46d5-a00d-eaba7019b332","created_by_photo_url":null,"childrens":null},{"id":"6640c44e-acbe-450f-b1cb-400469a374a5","text":"13","file_url":null,"file_name":null,"file_description":null,"file_size":null,"file_type":null,"file_mime_type":null,"parent_uuid":"89830116-f497-45fd-bbc3-fdda5ee7d6b7","created_at":"2024-05-13T03:19:00.856963Z","updated_at":"2024-05-13T03:19:00.856963Z","created_by_name":"user 50","created_by_id":"e6772dd5-9c96-46d5-a00d-eaba7019b332","created_by_photo_url":null,"childrens":null},{"id":"ca56037e-6b6d-4f79-aa83-a8b02b2f36f5","text":"12","file_url":null,"file_name":null,"file_description":null,"file_size":null,"file_type":null,"file_mime_type":null,"parent_uuid":"89830116-f497-45fd-bbc3-fdda5ee7d6b7","created_at":"2024-05-13T03:18:57.210173Z","updated_at":"2024-05-13T03:18:57.210173Z","created_by_name":"user 50","created_by_id":"e6772dd5-9c96-46d5-a00d-eaba7019b332","created_by_photo_url":null,"childrens":null}]
`

type ListKpiNote struct {
	UUID              string        `db:"uuid" json:"id"`
	Text              any           `db:"text" json:"text"`
	FileUrl           any           `db:"file_url" json:"file_url"`
	FileName          any           `db:"file_name" json:"file_name"`
	FileDescription   any           `db:"file_description" json:"file_description"`
	FileSize          any           `db:"file_size" json:"file_size"`
	FileType          any           `db:"file_type" json:"file_type"`
	FileMimeType      any           `db:"mime_type" json:"file_mime_type"`
	ParentUUID        string        `db:"parent_uuid" json:"parent_uuid"`
	CreatedAt         any           `db:"created_at" json:"created_at"`
	UpdatedAt         any           `db:"updated_at" json:"updated_at"`
	CreatedByName     any           `db:"created_by_name" json:"created_by_name"`
	CreatedByUUID     any           `db:"created_by" json:"created_by_id"`
	CreatedByPhotoUrl any           `db:"photo_url" json:"created_by_photo_url"`
	TotalRows         int           `db:"total_rows" json:"-"`
	Children          []ListKpiNote `json:"childrens"`
}

func main() {
	var allNotes []ListKpiNote
	if err := json.Unmarshal([]byte(dataStr), &allNotes); err != nil {
		log.Fatalln(err)
	}

	noteMap := make(map[string][]ListKpiNote)
	for _, note := range allNotes {
		noteMap[note.ParentUUID] = append(noteMap[note.ParentUUID], note)
	}

	var buildHierarchy func(parentUUID string) []ListKpiNote
	buildHierarchy = func(parentUUID string) []ListKpiNote {
		var hierarchy []ListKpiNote
		for _, child := range noteMap[parentUUID] {
			child.Children = buildHierarchy(child.UUID)
			hierarchy = append(hierarchy, child)
		}
		return hierarchy
	}

	data := buildHierarchy("")
	lala, _ := json.Marshal(data)
	fmt.Println(string(lala))
}
