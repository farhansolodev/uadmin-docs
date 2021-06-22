uadmin.F
========
`Back to Customizing Records Page`_

.. _Back to Customizing Records Page: https://uadmin-docs.readthedocs.io/en/latest/api/customizing_records.html

F is a field.

Structure:

.. code-block:: go

    type F struct {
        Name              string
        DisplayName       string
        ColumnName        string
        Type              string
        TypeName          string
        Value             interface{}
        Help              string
        Max               interface{}
        Min               interface{}
        Format            string
        DefaultValue      string
        Required          bool
        Pattern           string
        PatternMsg        string
        Hidden            bool
        ReadOnly          string
        Searchable        bool
        Filter            bool
        ListDisplay       bool
        FormDisplay       bool
        CategoricalFilter bool
        Translations      []translation
        Choices           []Choice
        IsMethod          bool
        ErrMsg            string
        ProgressBar       map[float64]string
        LimitChoicesTo    func(interface{}, *User) []Choice
        UploadTo          string
        Encrypt           bool
        Approval          bool
        NewValue          interface{}
        OldValue          interface{}
        ChangedBy         string
        ChangeDate        *time.Time
        ApprovalAction    ApprovalAction
        ApprovalDate      *time.Time
        ApprovalBy        string
        ApprovalID        uint
        Stringer          bool
    }

Parameters:

* **Name** - The name of the field
* **DisplayName** - The name that you want to display in the model. It is an alias name.
* **ColumnName** - The column name of the field
* **Type** - The field type (e.g. file, list, progress_bar)
* **TypeName** - The data type of the field (e.g. string, int, float64)
* **Value** - The value that you want to assign in a field
* **Help** - An instruction given to understand more details about the field or how to assign a value in a field
* **Max** - The maximum value the user can assign. It is applicable for numeric characters.
* **Min** - The minimum value the user can assign. It is applicable for numeric characters.
* **Format** - Implements formatted I/O with functions (e.g. %s - string, %d - Integer)
* **DefaultValue** - A value assigned automatically if you want to add a new record
* **Required** - A field that user must perform the given task(s). It cannot be skipped or left empty.
* **Pattern** - A regular expression
* **PatternMsg** - An error message if the user assigns a value that did not match the requested format
* **Hidden** - A feature to hide the component in the editing section of the form
* **ReadOnly** - A field that cannot be modified
* **Searchable** - A feature that allows the user to search for a field or column name
* **Filter** - A feature that allows the user to filter the record assigned in a model
* **ListDisplay** - A feature that will hide the field in the viewing section of the model if the value returns false
* **FormDisplay** - A feature that will hide the field in the editing section of the model if the value returns false
* **CategoricalFilter** - A feature that allows the user to filter the record assigned in a model in the form of combo box
* **Translations** - For multilingual fields
* **Choices** - A struct for the list of choices
* **IsMethod** - Check if the method should be included in the field list
* **ErrMsg** - An error message displayed beneath the input field
* **ProgressBar** - A feature used to measure the progress of the activity
* **LimitChoicesTo** - A feature used to append the fetched records in the drop down list
* **UploadTo** - A path where to save the uploaded files
* **Encrypt** - A feature used to encrypt the value in the database
* **Approval** - A feature used to set an approval permission in the field
* **NewValue** - A value that you want to replace from the old value
* **OldValue** - A value that was assigned before
* **ChangedBy** - Returns the username who changed the value of the field record
* **ChangeDate** - The date when the value of the field record was changed
* **ApprovalAction** - A selection of approval actions. There are two selections: Approved and Declined.
* **ApprovalDate** - The date when the value of the field record was approved
* **ApprovalBy** - Returns the username who approved the value of the field record
* **ApprovalID** - Returns the user ID who approved the value of the field record
* **Stringer** - A feature that assigns a field as a unique type

There are 2 ways you can do for initialization process using this function: one-by-one and by group.

One-by-one initialization:

.. code-block:: go

    func main(){
        // Some codes
        f := uadmin.F{}
        f.Name = "Name"
        f.DisplayName = "DisplayName"
        f.ColumnName = "column_name"
        f.Type = "Type"
        f.Value = "Value"
    }

By group initialization:

.. code-block:: go

    func main(){
        // Some codes
        f := uadmin.F{
            Name:        "Name",
            DisplayName: "DisplayName",
            ColumnName: "name",
            Type:        "Type",
            Value:       "Value",
        }
    }

In the following examples, we will use "by group" initialization process.

* `Example #1: String Data Type`_
* `Example #2: Progress Bar`_
* `Example #3: Choices`_
* `Example #4: Upload To`_
* `Example #5: Approval`_

**Example #1:** String Data Type
--------------------------------
`Back to Top`_

In this example, we will discuss the following:

* Change the actual name that is different from the name initialized in the model struct
* Assign a value automatically if you are creating a new record
* Assign a help note to instruct the user what to do on that field
* Set the field as Required
* Set a pattern and error message where only letters are acceptable
* Hide the field in the form and list
* Disable editing in a field
* Allows searching and filtering record(s)
* Adding an error message tag
* Encrypt the value of the name in the database

String Data Type (Links)
^^^^^^^^^^^^^^^^^^^^^^^^
* `Name Field in Category Model - DisplayName`_
* `Name Field in Category Model - DefaultValue`_
* `Name Field in Category Model - Help`_
* `Name Field in Category Model - Required`_
* `Name Field in Category Model - Pattern and PatternMsg`_
* `Name Field in Category Model - Searchable`_
* `Name Field in Category Model - Filter and CategoricalFilter`_
* `Name Field in Category Model - ReadOnly`_
* `Name Field in Category Model - ErrMsg`_
* `Name Field in Category Model - ListDisplay`_
* `Name Field in Category Model - FormDisplay`_
* `Name Field in Category Model - Hidden`_
* `Name Field in Category Model - Encrypt`_

Suppose you have this field in the Category model that has a primary key of 1 as shown below:

.. image:: ../assets/categorynamedefault.png
   :align: center

|

Go to main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            Name:     "Name",
            TypeName: "string", // Data Type
            Format:   "%s",     // String
        }

        // Model schema configurations
        modelschema := uadmin.ModelSchema{
            Name:      "Category", // Model name
            ModelName: "category", // URL
            ModelID:   1,          // Primary key
            Fields:    []uadmin.F{field},
        }

        // Call the field variable assigned in an array
        // Fields[0] = field
        modelschemafield := modelschema.Fields[0]

        // Call the schema of "category" where the field name is "Name"
        // modelschema.ModelName = "category"
        // modelschemafield.Name = "Name"
        category := uadmin.Schema[modelschema.ModelName].FieldByName(modelschemafield.Name)

        // Assign the TypeName field value to the Name field in Category model
        category.TypeName = modelschemafield.TypeName

        // Assign the Format field value to the Name field in Category model
        category.Format = modelschemafield.Format
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, we got the same result.

.. image:: ../assets/categorynamedefault.png
   :align: center

**Name Field in Category Model - DisplayName**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #1)`_

Now let's replace the actual field name. In order to do that, declare a value in the DisplayName field inside uadmin.F function then assign that value to the DisplayName field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            DisplayName: "Category Name",
        }

        // Some codes

        // Assign the DisplayName field value to the Name field in Category
        // model
        category.DisplayName = modelschemafield.DisplayName
    }

Run your application and go to Category model. As expected, the name has changed to "CATEGORY NAME".

.. image:: ../assets/categorydisplayname.png

**Name Field in Category Model - DefaultValue**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #1)`_

You can also assign a value in the Name field automatically when creating a new record. In order to do that, declare a value in the DefaultValue field inside uadmin.F function then assign that value to the DefaultValue field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            DefaultValue: "Type here",
        }

        // Some codes

        // Assign the DefaultValue field value to the Name field in Category
        // model
        category.DefaultValue = modelschemafield.DefaultValue
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, "Type here" value has assigned automatically in the Name field.

.. image:: ../assets/categorydefaultvalue.png
   :align: center

**Name Field in Category Model - Help**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #1)`_

Let's assign a help note in the Name field to instruct the user what to do on that field. In order to do that, declare a value in the Help field inside uadmin.F function then assign that value to the Help field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            Help: "Input a category name for your Todo List.",
        }

        // Some codes

        // Assign the Help field value to the Name field in Category model
        category.Help = modelschemafield.Help
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, the help note was displayed below the input Name field.

.. image:: ../assets/categorynamehelp.png
   :align: center

**Name Field in Category Model - Required**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #1)`_

Let's set a feature where the user needs to fill up the Name field. If the value is empty, the user will prompt the user that the value of the Name field should be assigned. In order to do that, declare a value in the Required field inside uadmin.F function then assign that value to the Required field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            Required: true,
        }

        // Some codes

        // Assign the Required field value to the Name field in Category
        // model
        category.Required = modelschemafield.Required
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. If you notice, there is an asterisk (\*) symbol located on the top right after the "Name:". Let's leave the Name field value as it is. If you click Save, the system will prompt the user that the Name must be filled out.

.. image:: ../assets/categorynamerequired.png
   :align: center

**Name Field in Category Model - Pattern and PatternMsg**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #1)`_

Let's set a feature where the user can assign letters only in the Name field. In order to do that, declare a value in the Required field inside uadmin.F function then assign that value to the Required field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            Pattern:      "^[a-zA-Z _]*$",
            PatternMsg:   "Your input must be a letter.",
        }

        // Some codes

        // Assign the Pattern field value to the Name field in Category model
        category.Pattern = modelschemafield.Pattern

        // Assign the PatternMsg field value to the Name field in Category
        // model
        category.PatternMsg = modelschemafield.PatternMsg
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. Let's assign a numeric value in the Name field. If you click Save, the system will prompt the user the the value of the Name field must assign letters only.

.. image:: ../assets/categorynamepattern.png
   :align: center

**Name Field in Category Model - Searchable**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #1)`_

Let's set a feature where the user can search the name in the Category model. In order to do that, declare a value in the Searchable field inside uadmin.F function then assign that value to the Searchable field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            Searchable:   true,
        }

        // Some codes

        // Assign the Searchable field value to the Name field in Category model
        category.Searchable = modelschemafield.Searchable
    }

Run your application and go to the Category model. As expected, there is a search engine at the top of the model form. Suppose you have two records as shown below:

.. image:: ../assets/categorysearchable.png

|

Let's search "Work" and see what happens.

.. image:: ../assets/categorysearchablework.png

**Name Field in Category Model - Filter and CategoricalFilter**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #1)`_

Let's set a feature where the user can filter the name in the Category model. In order to do that, declare a value in the Filter and CategoricalFilter fields inside uadmin.F function then assign that value to those fields in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            Filter:            true,
            CategoricalFilter: true,
        }

        // Some codes

        // Assign the Filter field value to the Name field in Category model
        category.Filter = modelschemafield.Filter

        // Assign the CategoricalFilter field value to the Name field in
        // Category model
        category.CategoricalFilter = modelschemafield.CategoricalFilter
    }

Run your application and go to the Category model. As expected, the combo box form highlighted on the right side is the CategoricalFilter to notify the user that the Category Name is the field that will be filtered. Now click the Filter button. Suppose you have two records as shown below:

.. image:: ../assets/categoryfilter.png

|

Assign "Work" in the Category Name. Click Filter button on the bottom right corner of the modal and see what happens.

.. image:: ../assets/categoryfilterwork.png

|

As expected, the Category record has filtered out where the name contains "Work".

.. image:: ../assets/categorynamefilterresult.png

**Name Field in Category Model - ReadOnly**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #1)`_

Let's set a feature where the user cannot modify a Name field in the Category model. In order to do that, declare a value in the ReadOnly field inside uadmin.F function then assign that value to the ReadOnly field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            ReadOnly: "true",
        }

        // Some codes

        // Assign the ReadOnly field value to the Name field in Category model
        category.ReadOnly = modelschemafield.ReadOnly
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, the Name field is now Read Only that means the value cannot be modified.

.. image:: ../assets/categorynamereadonly.png
   :align: center

**Name Field in Category Model - ErrMsg**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #1)`_

Let's set a feature where an error message will be displayed beneath the input Name field. In order to do that, declare a value in the ErrMsg field inside uadmin.F function then assign that value to the ErrMsg field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            ErrMsg: "This field cannot be modified.",
        }

        // Some codes

        // Assign the ErrMsg field value to the Name field in Category model
        category.ErrMsg = modelschemafield.ErrMsg
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, the error message was displayed beneath the input Name field.

.. image:: ../assets/categorynameerrmsg.png
   :align: center

**Name Field in Category Model - ListDisplay**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #1)`_

Let's set a feature that will hide the field or column name in the viewing section of the Category model. In order to do that, declare a value in the ListDisplay field inside uadmin.F function then assign that value to the ListDisplay field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            ListDisplay: false,
        }

        // Some codes

        // Assign the ListDisplay field value to the Name field in Category
        // model
        category.ListDisplay = modelschemafield.ListDisplay
    }

Run your application and go to the Category model. As expected, the Name Field in Category Model is now invisible in the list.

.. image:: ../assets/categorynamelistdisplay.png

**Name Field in Category Model - FormDisplay**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #1)`_

Let's set a feature that will hide the field in the editing section of the Category model. In order to do that, declare a value in the FormDisplay field inside uadmin.F function then assign that value to the FormDisplay field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            FormDisplay: false,
        }

        // Some codes

        // Assign the FormDisplay field value to the Name field in Category
        // model
        category.FormDisplay = modelschemafield.FormDisplay
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, the Name Field is now invisible in the Category model.

.. image:: ../assets/categorynameformdisplay.png
   :align: center

**Name Field in Category Model - Hidden**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #1)`_

Unlike in FormDisplay, the field will hide if the value is true. In order to hide the Name field in the Category model, declare a value in the Hidden field inside uadmin.F function then assign that value to the Hidden field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            Hidden: true,
        }

        // Some codes

        // Assign the Hidden field value to the Name field in Category model
        category.Hidden = modelschemafield.Hidden
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, the Name Field is now invisible in the Category model.

.. image:: ../assets/categorynameformdisplay.png
   :align: center

**Name Field in Category Model - Encrypt**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #1)`_

.. _Back to Top (Example #1): https://uadmin-docs.readthedocs.io/en/latest/api/customizing-records/f.html#example-1-string-data-type

Suppose you have two records as shown below:

.. image:: ../assets/categorynametworecords.png

|

Let's encrypt the value of the Name field in the Category Model. In order to do that, declare a value in the Encrypt field inside uadmin.F function then assign that value to the Encrypt field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            Encrypt: true,
        }

        // Some codes

        // Assign the Encrypt field value to the Name field in Category model
        category.Encrypt = modelschemafield.Encrypt
    }

Run your application. From your project folder, open uadmin.db with DB Browser for SQLite.

.. image:: ../assets/uadmindbsqlite.png
   :align: center

|

Click on Execute SQL.

.. image:: ../assets/executesqlhighlighted.png
   :align: center

|

Get all records by typing this command: **SELECT \* FROM categories** then click the right arrow icon to execute your SQL command.

.. image:: ../assets/selectfromcategories.png
   :align: center

|

As expected, the Name value is encrypted in the database.

.. image:: ../assets/categorynameencrypt.png
   :align: center

**Example #2:** Progress Bar
^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top`_

In this example, we will discuss the following:

* Change the actual name that is different from the name initialized in the model struct
* Set the type of the input field
* Set the color and maximum value in the progress bar
* Set the minimum and maximum limit the user can assign in the progress bar
* Assign a value automatically if you are creating a new record

Progress Bar (Links)
^^^^^^^^^^^^^^^^^^^^
* `Progress Field in Todo Model - DisplayName`_
* `Progress Field in Todo Model - Type`_
* `Progress Field in Todo Model - ProgressBar`_
* `Progress Field in Todo Model - Max and Min`_
* `Progress Field in Todo Model - DefaultValue`_

Suppose you have this field in the Todo model that has a primary key of 1 as shown below:

.. image:: ../assets/todoprogressdefault.png

|

Go to main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            Name:     "Progress",
            TypeName: "float64", // Data Type
            Format:   "%d",      // Integer
        }

        // Model schema configurations
        modelschema := uadmin.ModelSchema{
            Name:      "Todo", // Model name
            ModelName: "todo", // URL
            ModelID:   1,      // Primary key
            Fields:    []uadmin.F{field},
        }

        // Call the field variable assigned in an array
        // Fields[0] = field
        modelschemafield := modelschema.Fields[0]

        // Call the schema of "todo" where the field name is "Progress"
        // modelschema.ModelName = "todo"
        // modelschemafield.Name = "Progress"
        todo := uadmin.Schema[modelschema.ModelName].FieldByName(modelschemafield.Name)

        // Assign the TypeName field value to the Progress field in Todo model
        todo.TypeName = modelschemafield.TypeName
    }

Run your application and go to the Todo model. As expected, we got the same result.

.. image:: ../assets/todoprogressdefault.png

**Progress Field in Todo Model - DisplayName**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #2)`_

Now let's replace the actual field name. In order to do that, declare a value in the DisplayName field inside uadmin.F function then assign that value to the DisplayName field in the Todo model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            DisplayName: "Current Progress",
        }

        // Some codes

        // Assign the DisplayName field value to the Progress field in Todo
        // model
        todo.DisplayName = modelschemafield.DisplayName
    }

Run your application and go to the Todo model. As expected, the name has changed to "CURRENT PROGRESS".

.. image:: ../assets/todoprogressdisplayname.png

**Progress Field in Todo Model - Type**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #2)`_

Let's convert the input type to the progress bar. In order to do that, declare a value in the Type field inside uadmin.F function then assign that value to the Type field in the Todo model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            Type: "progress_bar",
        }

        // Some codes

        // Assign the Type field value to the Progress field in Todo model
        todo.Type = modelschemafield.Type
    }

Run your application and go to the Todo model. As expected, the field type has changed from regular to a progress bar. However, the appearance does not look good because we have not assigned the value and color of the progress bar yet.

.. image:: ../assets/todoprogresstype.png

**Progress Field in Todo Model - ProgressBar**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #2)`_

Let's assign the value and the color of the progress bar. In order to do that, declare a value in the ProgressBar field inside uadmin.F function then assign that value to the ProgressBar field in the Todo model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes

            // 100.0 - maximum value
            // #07c - blue color
            ProgressBar: map[float64]string{100.0: "#07c"},
        }

        // Some codes

        // Assign the ProgressBar field value to the Progress field in Todo
        // model
        todo.ProgressBar = modelschemafield.ProgressBar
    }

Run your application and go to the Todo model. As expected, the appearance of the progress bar is now good enough.

.. image:: ../assets/todoprogressbar.png

**Progress Field in Todo Model - Max and Min**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #2)`_

Let's set a limitation where the user can assign a value between the range of 0 and 100. In order to do that, declare a value in the Max and Min fields inside uadmin.F function then assign that value to those fields in the Todo model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            Max: "100",
            Min: "0",
        }

        // Some codes

        // Assign the Max field value to the Progress field in Todo model
        todo.Max = modelschemafield.Max

        // Assign the Min field value to the Progress field in Todo model
        todo.Min = modelschemafield.Min
    }

Run your application and go to the Todo model. Let's put a numeric value outside the range of 0 and 100 in the Progress field and see what happens.

.. image:: ../assets/todoprogressmax.png

**Progress Field in Todo Model - DefaultValue**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #2)`_

.. _Back to Top (Example #2): https://uadmin-docs.readthedocs.io/en/latest/api/customizing-records/f.html#example-2-progress-bar

You can also assign a value in the Progress field automatically when creating a new record. In order to do that, declare a value in the DefaultValue field inside uadmin.F function then assign that value to the DefaultValue field in the Todo model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            DefaultValue: "50",
        }

        // Some codes

        // Assign the DefaultValue field value to the Progress field in Todo
        // model
        todo.DefaultValue = modelschemafield.DefaultValue
    }

Run your application, go to the Todo model and click Add New Todo button on the top right corner of the screen. As expected, "50" value has assigned automatically in the Progress field.

.. image:: ../assets/todoprogressdefaultvalue.png

**Example #3:** Choices
^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top`_

In this example, we will discuss the following:

* Build choices
* Hide the field in the form and list

Choices (Links)
^^^^^^^^^^^^^^^
* `Nationality Field in Friend Model - Choices`_
* `Nationality Field in Friend Model - ListDisplay`_
* `Nationality Field in Friend Model - FormDisplay`_
* `Nationality Field in Friend Model - Hidden`_

Suppose you have the given source code in friend.go where Nationality is the type of the drop down list:

.. code-block:: go

    // Nationality ...
    type Nationality int

    // Chinese ...
    func (Nationality) Chinese() Nationality {
        return 1
    }

    // Filipino ...
    func (Nationality) Filipino() Nationality {
        return 2
    }

    // Others ...
    func (Nationality) Others() Nationality {
        return 3
    }

And you have this field in the Friend model that has a primary key of 1 containing three choices which are Chinese, Filipino, and Others as shown below:

.. image:: ../assets/friendnationalitydefault.png

|

Go to main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            Name:     "Nationality",
            Type:     "list",
            TypeName: "Nationality",
        }

        // Model schema configurations
        modelschema := uadmin.ModelSchema{
            Name:      "Friend", // Model name
            ModelName: "friend", // URL
            ModelID:   1,        // Primary key
            Fields:    []uadmin.F{field},
        }

        // Call the field variable assigned in an array
        // Fields[0] = field
        modelschemafield := modelschema.Fields[0]

        // Call the schema of "friend" where the field name is "Nationality"
        // modelschema.ModelName = "friend"
        // modelschemafield.Name = "Nationality"
        friend := uadmin.Schema[modelschema.ModelName].FieldByName(modelschemafield.Name)

        // Assign the Type field value to the Nationality field in Friend model
        friend.Type = modelschemafield.Type

        // Assign the TypeName field value to the Nationality field in Friend
        // model
        friend.TypeName = modelschemafield.TypeName
    }

Run your application, go to the Friend model and click Add New Friend button on the top right corner of the screen. As expected, we got the same result.

.. image:: ../assets/friendnationalitydefault.png

**Nationality Field in Friend Model - Choices**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #3)`_

Let's build a choice that includes Chinese and Filipino and excludes Others. In order to do that, declare a value in the Choices field inside uadmin.F function then assign that value to the Choices field in the Friend model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes

            // K is the ID of the choice
            // V is the value of the choice
            Choices: []uadmin.Choice{
                {K: 0, V: " - "},
                {K: 1, V: "Chinese"},
                {K: 2, V: "Filipino"},
            },
        }

        // Some codes

        // Assign the Choices field value to the Nationality field in Friend
        // model
        friend.Choices = modelschemafield.Choices
    }

Run your application, go to the Friend model and click Add New Friend button on the top right corner of the screen. As expected, Chinese and Filipino choices are included in the list.

.. image:: ../assets/friendnationalitychoices.png

**Nationality Field in Friend Model - ListDisplay**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #3)`_

Let's set a feature that will hide the field or column name in the viewing section of the Friend model. In order to do that, declare a value in the ListDisplay field inside uadmin.F function then assign that value to the ListDisplay field in the Friend model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            ListDisplay: false,
        }

        // Some codes

        // Assign the ListDisplay field value to the Nationality field in
        // Friend model
        friend.ListDisplay = modelschemafield.ListDisplay
    }

Run your application and go to the Friend model. As expected, the Nationality Field in the Friend Model is now invisible in the list.

.. image:: ../assets/friendnationalitylistdisplay.png

**Nationality Field in Friend Model - FormDisplay**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #3)`_

Let's set a feature that will hide the field in the editing section of the Friend model. In order to do that, declare a value in the FormDisplay field inside uadmin.F function then assign that value to the FormDisplay field in the Friend model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            FormDisplay: false,
        }

        // Some codes

        // Assign the FormDisplay field value to the Nationality field in
        // Friend model
        friend.FormDisplay = modelschemafield.FormDisplay
    }

Run your application, go to the Friend model and click Add New Friend button on the top right corner of the screen. As expected, the Nationality Field is now invisible in the Friend model.

.. image:: ../assets/friendnationalityformdisplay.png

**Nationality Field in Friend Model - Hidden**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #3)`_

.. _Back to Top (Example #3): https://uadmin-docs.readthedocs.io/en/latest/api/customizing-records/f.html#example-3-choices

Unlike in FormDisplay, the field will hide if the value is true. In order to hide the Name field in the Category model, declare a value in the Hidden field inside uadmin.F function then assign that value to the Hidden field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            Hidden: true,
        }

        // Some codes

        // Assign the Hidden field value to the Nationality field in Friend
        // model
        friend.Hidden = modelschemafield.Hidden
    }

Run your application, go to the Friend model and click Add New Friend button on the top right corner of the screen. As expected, the Nationality Field is now invisible in the Friend model.

.. image:: ../assets/friendnationalityformdisplay.png

**Example #4:** Upload To
^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top`_

In this example, we will discuss the following:

* Set the type of the input field
* Set the field as Required
* Set a path where to upload the files
* Adding an error message tag
* Hide the field in the form and list

Upload To (Links)
^^^^^^^^^^^^^^^^^
* `File Field in Category Model - Type`_
* `File Field in Category Model - Required`_
* `File Field in Category Model - UploadTo`_
* `File Field in Category Model - ErrMsg`_
* `File Field in Category Model - ListDisplay`_
* `File Field in Category Model - FormDisplay`_
* `File Field in Category Model - Hidden`_

Suppose you have this field in the Category model that has a primary key of 1 as shown below:

.. image:: ../assets/categoryfiledefault.png
   :align: center

|

Go to main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            Name:     "File",
            TypeName: "string",
        }

        // Model schema configurations
        modelschema := uadmin.ModelSchema{
            Name:      "Category", // Model name
            ModelName: "category", // URL
            ModelID:   1,          // Primary key
            Fields:    []uadmin.F{field},
        }

        // Call the field variable assigned in an array
        // Fields[0] = field
        modelschemafield := modelschema.Fields[0]

        // Call the schema of "category" where the field name is "File"
        // modelschema.ModelName = "category"
        // modelschemafield.Name = "File"
        category := uadmin.Schema[modelschema.ModelName].FieldByName(modelschemafield.Name)

        // Assign the TypeName field value to the File field in Category model
        category.TypeName = modelschemafield.TypeName
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, we got the same result.

.. image:: ../assets/categoryfiledefault.png
   :align: center

**File Field in Category Model - Type**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #4)`_

Let’s convert the input type to the file. In order to do that, declare a value in the Type field inside uadmin.F function then assign that value to the Type field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            Type: "file",
        }

        // Some codes

        // Assign the Type field value to the File field in Category model
        category.Type = modelschemafield.Type
    }

Run your application and go to the Category model. As expected, the field type has changed from regular to a file input. 

.. image:: ../assets/categoryfiletype.png
   :align: center

**File Field in Category Model - Required**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #4)`_

Let’s set a feature where the user needs to fill up the File field. If the value is empty, the user will prompt the user that the value of the File field should be assigned. In order to do that, declare a value in the Required field inside uadmin.F function then assign that value to the Required field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            Required: true,
        }

        // Some codes

        // Assign the Required field value to the File field in Category model
        category.Required = modelschemafield.Required
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. If you notice, there is an asterisk (\*) symbol located on the top right after the "File:". Let's leave the File field value as it is. If you click Save, nothing will happen until you fill out the File field.

.. image:: ../assets/categoryfilerequired.png
   :align: center

**File Field in Category Model - UploadTo**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #4)`_

Let's set a feature where the uploaded file will save in the specified path on your project folder. In order to do that, declare a value in the UploadTo field inside uadmin.F function then assign that value to the UploadTo field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            UploadTo: "/media/files/",
        }

        // Some codes

        // Assign the UploadTo field value to the File field in Category model
        category.UploadTo = modelschemafield.UploadTo
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. Let's add a new record that includes the uploaded file from your computer (e.g. Windows Installation.pdf).

.. image:: ../assets/categoryinstallationrecord.png
   :align: center

|

Result:

.. image:: ../assets/categoryinstallationrecordresult.png

|

From your project folder, go to /media/files/(generated_folder_name)/. As expected, the "Windows Installation.pdf" file was saved on that path.

.. image:: ../assets/categoryfileuploadto.png
   :align: center

**File Field in Category Model - ErrMsg**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #4)`_

Let’s set a feature where an error message will be displayed beneath the input File field. In order to do that, declare a value in the ErrMsg field inside uadmin.F function then assign that value to the ErrMsg field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            ErrMsg: "Invalid format",
        }

        // Some codes

        // Assign the ErrMsg field value to the File field in Category model
        category.ErrMsg = modelschemafield.ErrMsg
    }

Run your application, go to the Category model and click the existing record that you have. As expected, the error message was displayed beneath the input File field.

.. image:: ../assets/categoryfileerrmsg.png

**File Field in Category Model - ListDisplay**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #4)`_

Let’s set a feature that will hide the field or column name in the viewing section of the Category model. In order to do that, declare a value in the ListDisplay field inside uadmin.F function then assign that value to the ListDisplay field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            ListDisplay: false,
        }

        // Some codes

        // Assign the ListDisplay field value to the File field in Category
        // model
        category.ListDisplay = modelschemafield.ListDisplay
    }

Run your application and go to the Category model. As expected, the File Field in the Category Model is now invisible in the list.

.. image:: ../assets/categoryfilelistdisplay.png

**File Field in Category Model - FormDisplay**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #4)`_

Let’s set a feature that will hide the field in the editing section of the Category model. In order to do that, declare a value in the FormDisplay field inside uadmin.F function then assign that value to the FormDisplay field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            FormDisplay: false,
        }

        // Some codes

        // Assign the FormDisplay field value to the File field in Category
        // model
        category.FormDisplay = modelschemafield.FormDisplay
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, the File Field is now invisible in the Category model.

.. image:: ../assets/categoryfileformdisplay.png
   :align: center

**File Field in Category Model - Hidden**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top (Example #4)`_

.. _Back to Top (Example #4): https://uadmin-docs.readthedocs.io/en/latest/api/customizing-records/f.html#example-4-upload-to

Unlike in FormDisplay, the field will hide if the value is true. In order to hide the File field in the Category model, declare a value in the Hidden field inside uadmin.F function then assign that value to the Hidden field in the Category model.

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            // Some codes
            Hidden: true,
        }

        // Some codes

        // Assign the Hidden field value to the File field in Category model
        category.Hidden = modelschemafield.Hidden
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, the File Field is now invisible in the Category model.

.. image:: ../assets/categoryfileformdisplay.png
   :align: center

**Example #5:** Approval
^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. _Back to Top: https://uadmin-docs.readthedocs.io/en/latest/api/customizing-records/f.html#uadmin-f

In this example, we will discuss on how to set an approval permission in the field.

Suppose you have "johndoe" account in the User model that is not an admin.

.. image:: ../assets/johndoerecord.png

|

User permissions are set for that user in Approvals and Friends models without an approval access.

.. image:: ../assets/johndoeuserpermissionapproval.png

|

And you have the Name field in the Friend model that has a primary key of 1 as shown below:

.. image:: ../assets/friendnamedefault.png
   :align: center

|

Go to main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Some codes

        // Field configurations
        field := uadmin.F{
            Name:     "Name",
            TypeName: "string", // Data Type
            Format:   "%s",     // String
            Approval: true,
        }

        // Model schema configurations
        modelschema := uadmin.ModelSchema{
            Name:      "Friend", // Model name
            ModelName: "friend", // URL
            ModelID:   1,        // Primary key
            Fields:    []uadmin.F{field},
        }

        // Call the field variable assigned in an array
        // Fields[0] = field
        modelschemafield := modelschema.Fields[0]

        // Call the schema of "friend" where the field name is "Name"
        // modelschema.ModelName = "friend"
        // modelschemafield.Name = "Name"
        friend := uadmin.Schema[modelschema.ModelName].FieldByName(modelschemafield.Name)

        // Assign the TypeName field value to the Name field in Friend model
        friend.TypeName = modelschemafield.TypeName

        // Assign the Format field value to the Name field in Friend model
        friend.Format = modelschemafield.Format

        // Assign the Approval field value to the Name field in Friend model
        friend.Approval = modelschemafield.Approval
    }

Run your application and login your account using "johndoe".

.. image:: ../assets/johndoelogin.png
   :align: center

|

As you can see, only the Approvals and Friends models are accessible in the dashboard. It is based on the user permission that was assigned on that user. Now click on "FRIENDS".

.. image:: ../assets/friendsapprovalhighlighted.png

|

Click "Add New Friend" to create a new record.

.. image:: ../assets/addnewfriend.png

|

Let's input the following field for this record.

.. image:: ../assets/johndoefriendrecord.png
   :align: center

|

Result

.. image:: ../assets/johndoenameempty.png

|

Based on the result, the name does not show up because we need an approval to someone who has approval access. Now logout johndoe account then login an admin account.

.. image:: ../assets/loginformadmin.png
   :align: center

|

From uAdmin dashboard, go to the Friends model, click the record that you have created, and in the input box of the Name field, there is a yellow warning sign on the left side that means it needs an approval to someone who has approval access. Now click the highlighted area below.

.. image:: ../assets/johndoeapprovalbutton.png
   :align: center

|

The admin will review the record that was created by a "johndoe" user. If you think his record is satisfactory, choose Approved in Approval Action then click Save and Continue on the bottom right corner of the screen.

.. image:: ../assets/johndoeapprovalreview.png

|

It is shown that the one who approved the record is an admin with an approved date. Now click View Record button to see the result.

.. image:: ../assets/johndoeviewrecord.png
   :align: center

|

The input Name field has a checkmark sign that means the record created by "johndoe" was approved.

.. image:: ../assets/johndoeapprovedrecord.png
   :align: center
