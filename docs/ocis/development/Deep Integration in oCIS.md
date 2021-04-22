# Deep Integration in oCIS

These are thoughts about APIs for "deeper" integrations of extensions with oCIS. Deep means that the extension uses oCIS as data storage for meta data and binary data and integrates with the oCIS APIs to query metadata, versions and user information such as permissions.

## Rational

oCIS is the perfect scalable, easy to use and performant data storage for all kind of unstructured data. 

To underline that statement oCIS must be very easy to integrate with. For that we need to expose special APIs.

oCIS is easing development of extensions by providing a lot of "services" that an extension does not need to care for such as user authentication, autorization, abstraction of storage, but also things like clients on all platforms.

## Interfaces 

### Overview

This diagram provides an interface overview. 

![Interface Overview](integrations.svg) 

### User management

oCIS provides a the extension with an validated user, the group and permission settings.

### Sharing

oCIS provides sharing of Resources within the group of ownCloud users and groups (private Links), as well as external users by public links.

The extension handles view or editing of the shares.

### Events

oCIS notifies the extension of certain events, ie. new file is available or user has been added. The extension can start actions based on that, for example index the new file.

These asynchronous notifications are the base for workflow integration.

### Versions and Trash Bin

If a resource is going to be modified within the extension it can create a version before that happens, so that changes can be tracked through oCIS. 

### Blob Data - ie. Files

oCIS manages to store and deliver files. Abstraction of storages, remote access, as well as sharing and syncing.

### Meta Data

Meta data is stored in oCIS. For extension that means that no own metadata storage and retrieval needs to be written. Also, the extension can make use of more metadata than it created by itself, ie additional user metadata that was created by a different app.

The metadata is completely generic, ie. stored as tuple of name, data and type, such as for example "Name" - "Klaas" - "String" or "Birthday" - "23.04.2006" - "Date". The name includes a namespace. 

Lists of the similar Meta tags can be pushed by using the same name and type and different value.

For example if an image viewer identifies that an picture displays a kid, a cat and a dog, it stores three metadata like this:

| Name | Type | Value |
|----------|-----------|-----------------|
|myext:content|String|kid|
|myext:content|String|dog|
|myext:content|String|cat|


### Resource Query Engine

Allows to query lists of Resources (single files or file spaces) based on combinations of metadata. The exension uses it to get lists of files to work on which it gets access through the Blob Data API.

## Examples

This describes a few examples how extensions could interact with the oCIS core and what can be achieved.

### Image Viewer Displays a File

TBD

### Image Viewer turns a File

The extension loads the file via the Blob API and modifies the file.  When it is stored back via the Blob API, oCIS creates a new version of the file automatically. 

The extension can query the existing versions via the Versions API.

### Image Viewer lists all files with a Cat.

A user working in the extension looks for all images that diplay a cat.

The extension queries the Metadata for the list of resources that have the Metadata "myext::content contains cat"

### QA Tool for Medical Ampule Production

A software for quality assurance on medical ampules takes a photo of every produced ampule and runs a check on the image to verify that the glass is free of damage. The day production is two million ampules and for documentation reasons, the tool stores each image into oCIS. For every batch of 10,000 ampules it needs to provide a quality report.

#### **Storing of the Image**

After the image of the ampule was taken, the extension uses the Blob Data API to store the file into oCIS. In return it gets an UUID of the resource in oCIS that will be used in further operations. As the size of the images is very defined, it uses a synchronous API.

#### **Image Processing Results**

After the image was processed by the extension it uses the UUID to store the result (ie. Successfull or Failed) and a BatchID (identifying the current batch) using the Meta Data API. 

#### **Finalising the Batch**

When the batch of ampules is finished the Extension uses the Metadata API to query for a list of good samples and the amount of fails. Note, it does not retrieve the image files back, but only works on the meta data. With that data it generates a quality report.

#### **Views for the Quality Manager**

Since the Extension is running as part of the automation software of the production machine, it does not provide web access to the data it produced.

A simple ownCloud Web extension can provide the Quality Manager with a comfortable web based view of the quality reports, which she can simply view on her mobile device, and even load the images of the broken ampules for example.

The ongoing processing of the machine can be monitored in Grafana via the ownCloud Grafana App to be able to react on unexpected machine stops. For that it monitors the upload of the quality images.

#### **Archiving** 

Every night when the machine is  down for maintenance ownCloud archives the image data of the previous day from the expensive fast storage to a slower long term storage system.  
