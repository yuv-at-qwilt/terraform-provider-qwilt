
#Create an empty resource to import into.

#After the import is complete, manually set the required attributes
#in the resource based on the imported state.


resource "qwilt_cdn_certificate_template" "example" {
}

terraform import qwilt_cdn_certificate_template.example <certificate_template_id>