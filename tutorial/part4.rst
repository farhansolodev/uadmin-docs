uAdmin Tutorial Part 4 - Linking Models
=======================================
Linking a model to another model is as simple as creating a field using a foreign key. Foreign Key is the key used to link two models together.

**What is the purpose of the foreign key?** The purpose of the foreign key is to ensure referential integrity of the data. In other words, only values that are supposed to appear in the database are permitted.

In the example below we linked the Category model into Todo model, now the Todo model will return its data as a field in the Category model.

.. code-block:: go

    package models

    import (
        "time"

        "github.com/uadmin/uadmin"
    )

    // Todo !
    type Todo struct {
        uadmin.Model
        Name        string
        Description string `uadmin:"html"`
        Category    Category // <-- Category Model
        CategoryID  uint     // <-- Category ID
        TargetDate  time.Time
        Progress    int `uadmin:"progress_bar"`
    }

Result

.. image:: assets/categoryaddedintodo.png

|

Create a file named **friend.go** inside your models folder with the following codes below:

.. code-block:: go

    package models

    import (
        "github.com/uadmin/uadmin"
    )

    // Friend Model !
    type Friend struct {
        uadmin.Model
        Name     string `uadmin:"required"`
        Email    string `uadmin:"email"`
        Password string `uadmin:"password;list_exclude"`
    }

Now register the model on main.go where `models` is the package name and `Friend` is the model name:

.. code-block:: go

    func main() {
        uadmin.Register(
            models.Todo{},
            models.Category{},
            models.Friend{}, // <-- place it here
        )
        uadmin.StartServer()
    }

Set the foreign key of a Friend model to the Todo model and apply the tag "help" to show that who will be a part of your todo activity.

.. code-block:: go

    package models

    import (
        "time"

        "github.com/uadmin/uadmin"
    )

    // Todo Model !
    type Todo struct {
        uadmin.Model
        Name        string
        Description string `uadmin:"html"`
        Category    Category
        CategoryID  uint

        // Friend Model
        Friend      Friend `uadmin:"help:Who will be a part of your activity?"`
        // Friend ID
        FriendID    uint

        TargetDate  time.Time
        Progress    int `uadmin:"progress_bar"`
    }

As expected, the Friend model is added in the uAdmin Dashboard.

.. image:: assets/friendsmodelselected.png

|

Let's create a new data in the Friend model.

.. image:: assets/friendsdata.png

|

Result

.. image:: assets/friendsdataoutput.png

|

As you can see, the password field is not shown in the output. Why? If you go back to the Friend model, the password field has the tag name "list_exclude". It means it will hide the field or column name in the model structure.

Go back to Todo model and see what happens.

.. image:: assets/friendsaddedintodo.png

|

Congrats, now you know how to link a model using a foreign key.

Click `here`_ to view our progress so far.

In the `next part`_ we will talk about register inlines and how to create a drop down list in the field manually.

.. _here: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part4.html

.. _next part: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part5.html

.. toctree::
   :maxdepth: 1

   full_code/part4
