package timetask

import "testing"

const modules_xml = `<?xml version="1.0" encoding="UTF-8" ?>
<ajax-response>
	<response type="object" id="ModuleList">
					<item>
						<moduleid><![CDATA[55555]]></moduleid>
						<modulename><![CDATA[R&amp;D]]></modulename>
						</item>
					<item>
						<moduleid><![CDATA[77777]]></moduleid>
						<modulename><![CDATA[Sprint 1]]></modulename>
						</item>
					<item>
						<moduleid><![CDATA[222222]]></moduleid>
						<modulename><![CDATA[Sprint 2]]></modulename>
						</item>
			</response>
</ajax-response>`

func TestModuleParseXML(t *testing.T) {
	modules, err := ModuleParseXML(modules_xml)
	if err != nil {
		t.Error(err)
	}

	_ = []Module{ // wanted
		Module{
			ID:   55555,
			Name: "R&amp;D",
		},
		Module{
			ID:   77777,
			Name: "Sprint 1",
		},
		Module{
			ID:   222222,
			Name: "Sprint 2",
		},
	}

	// Need a way to compare slices
	// if modules != wanted {
	// 	t.Errorf("Module parsing failed. Wanted %+v got %+v", wanted, modules)
	// }

	t.Logf("%+v\n", modules)
}
