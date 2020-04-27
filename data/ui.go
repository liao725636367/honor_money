package data
//GetTpl 将glade程序模板放到函数里面方便跟程序一起打包
func GetTpl()string{
	//<object class="GtkApplicationWindow" id="window1">
	return  `<?xml version="1.0" encoding="UTF-8"?>
<interface>
  <!-- interface-requires gtk+ 3.0 -->
  <object class="GtkWindow" id="window1">
    <property name="width_request">300</property>
    <property name="height_request">200</property>
    <property name="can_focus">False</property>
    <property name="halign">start</property>
    <property name="valign">start</property>
    <child>
      <object class="GtkFixed" id="fixed1">
        <property name="visible">True</property>
        <property name="can_focus">False</property>
        <child>
          <object class="GtkEntry" id="entry1">
            <property name="width_request">156</property>
            <property name="height_request">30</property>
            <property name="visible">True</property>
            <property name="can_focus">True</property>
            <property name="halign">start</property>
            <property name="valign">start</property>
            <property name="invisible_char">●</property>
            <property name="input_purpose">number</property>
          </object>
          <packing>
            <property name="x">62</property>
            <property name="y">106</property>
          </packing>
        </child>
        <child>
          <object class="GtkLabel" id="label1">
            <property name="width_request">50</property>
            <property name="height_request">30</property>
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="label" translatable="yes">连续刷</property>
            <property name="single_line_mode">True</property>
            <property name="track_visited_links">False</property>
          </object>
          <packing>
            <property name="x">9</property>
            <property name="y">106</property>
          </packing>
        </child>
        <child>
          <object class="GtkLabel" id="label2">
            <property name="width_request">30</property>
            <property name="height_request">30</property>
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="label" translatable="yes">次</property>
            <property name="track_visited_links">False</property>
          </object>
          <packing>
            <property name="x">226</property>
            <property name="y">106</property>
          </packing>
        </child>
        <child>
          <object class="GtkButton" id="button1">
            <property name="label" translatable="yes">开始</property>
            <property name="width_request">50</property>
            <property name="height_request">30</property>
            <property name="visible">True</property>
            <property name="can_focus">True</property>
            <property name="receives_default">True</property>
            <property name="focus_on_click">False</property>
          </object>
          <packing>
            <property name="x">9</property>
            <property name="y">145</property>
          </packing>
        </child>
        <child>
          <object class="GtkCheckButton" id="checkbutton1">
            <property name="label" translatable="yes">刷完关机</property>
            <property name="width_request">100</property>
            <property name="height_request">30</property>
            <property name="can_focus">False</property>
            <property name="receives_default">False</property>
            <property name="halign">start</property>
            <property name="valign">start</property>
            <property name="focus_on_click">False</property>
            <property name="xalign">0</property>
            <property name="draw_indicator">True</property>
          </object>
          <packing>
            <property name="x">117</property>
            <property name="y">145</property>
          </packing>
        </child>
        <child>
          <object class="GtkButton" id="button2">
            <property name="label" translatable="yes">获取设备</property>
            <property name="width_request">70</property>
            <property name="height_request">30</property>
            <property name="visible">True</property>
            <property name="can_focus">True</property>
            <property name="receives_default">True</property>
          </object>
          <packing>
            <property name="x">196</property>
            <property name="y">18</property>
          </packing>
        </child>
        <child>
          <object class="GtkComboBoxText" id="comboboxtext1">
            <property name="width_request">180</property>
            <property name="height_request">30</property>
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="valign">center</property>
            <property name="entry_text_column">0</property>
            <property name="id_column">1</property>
            <items>
              <item translatable="yes">--请选择--</item>
            </items>
          </object>
          <packing>
            <property name="x">9</property>
            <property name="y">19</property>
          </packing>
        </child>
        <child>
          <object class="GtkLabel" id="label3">
            <property name="width_request">252</property>
            <property name="height_request">30</property>
            <property name="visible">True</property>
            <property name="can_focus">False</property>
          </object>
          <packing>
            <property name="x">11</property>
            <property name="y">58</property>
          </packing>
        </child>
      </object>
    </child>
  </object>
</interface>
`
}
