Document System Tutorial Part 1 - Build A Project
=================================================
In this part, we will cover on building and preparing a project from scratch.

First of all, let's create a folder for your project and prepare it.

.. code-block:: bash

    $ mkdir -p ~/go/src/github.com/your_name/document_system
    $ cd ~/go/src/github.com/your_name/document_system
    $ uadmin prepare
    [   OK   ]   Created: /home/pc_name/go/src/github.com/your_name/document_system/models
    [   OK   ]   Created: /home/pc_name/go/src/github.com/your_name/document_system/api
    [   OK   ]   Created: /home/pc_name/go/src/github.com/your_name/document_system/views
    [   OK   ]   Created: /home/pc_name/go/src/github.com/your_name/document_system/media
    [   OK   ]   Created: /home/pc_name/go/src/github.com/your_name/document_system/static
    [   OK   ]   Created: /home/pc_name/go/src/github.com/your_name/document_system/templates

Use your favorite editor to create "main.go” inside that path. Put the following code in "main.go”.

.. code-block:: go

    package main

    import (
        "github.com/uadmin/uadmin"
    )

    func main() {
        // Assign Site Name value as "Document System"
        // NOTE: This code works only if database does not exist yet.
        uadmin.SiteName = "Document System"

        // Activates a uAdmin server
        uadmin.StartServer()
    }

.. IMPORTANT::
   In Windows, you need to use localhost in order to run your application (e.g. http://localhost:8080). Another way is to set your loopback Internet protocol (IP) address by using uadmin.BindIP to establish an IP connection to the same machine or computer being used by the end-user.

Sample:

.. code-block:: go

    func main(){
        // Assign BindIP value to "127.0.0.1" in the Settings
        // NOTE: This code works only if database does not exist yet.
        uadmin.BindIP = "127.0.0.1"
    }

Now to run your code (Linux and Apple macOS):

.. code-block:: bash

    go build; ./document_system

In Windows:

.. code-block:: bash

    go build & document_system

Result:

.. code-block:: bash

    [   OK   ]   Initializing DB: [13/13]
    [   OK   ]   Initializing Languages: [185/185]
    [  INFO  ]   Auto generated admin user. Username: admin, Password: admin.
    [   OK   ]   Synching System Settings: [51/51]
    [   OK   ]   Server Started: http://0.0.0.0:8080
             ___       __          _
      __  __/   | ____/ /___ ___  (_)___
     / / / / /| |/ __  / __  __ \/ / __ \
    / /_/ / ___ / /_/ / / / / / / / / / /
    \__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/

Open your browser and type the IP address above. Then login using "admin” as username and password.

.. image:: assets/loginform.png
   :align: center

|

You will be greeted by the Document System dashboard that contains the system models built in uAdmin.

.. image:: assets/documentsystemdashboard.png

|

Click `here`_ to view our progress so far.

In the `next part`_, we will discuss about creating and registering an external models in our application.

.. _here: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part1.html
.. _next part: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/part2.html

.. toctree::
   :maxdepth: 1

   full_code/part1
