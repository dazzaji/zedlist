package query

import (
	"testing"

	"github.com/gernest/zedlist/modules/forms"

	"github.com/gernest/zedlist/models"
)

func TestJobQuery(t *testing.T) {
	sample := []*models.Job{
		{Title: "first"},
		{Title: "second"},
		{Title: "third"},
	}
	defer func() {
		err := Delete(&sample)
		if err != nil {
			t.Error(err)
		}
	}()

	// CreaateJob
	for _, v := range sample {
		err := CreateJob(v)
		if err != nil {
			t.Errorf("creating new job %v", err)
		}
	}

	// GetJobByID
	for _, v := range sample {
		j, err := GetJobByID(v.ID)
		if err != nil {
			t.Errorf("getting a job %v", err)
		}
		if j.Title != v.Title {
			t.Errorf("expected %s got %s", v.Title, j.Title)
		}
	}

	// GetAllJobs
	jobs, err := GetALLJobs()
	if err != nil {
		t.Errorf("getting all jobs %v", err)
	}
	if len(jobs) < len(sample) {
		t.Errorf("expecetd %d to be greater than %d", len(jobs), len(sample))
	}
	if jobs == nil {
		t.Error("exppected all jobs got nil intead")
	}

	// GetLatestJobs
	latest, err := GetLatestJobs()
	if err != nil {
		t.Errorf("getting latest jobs %v", err)
	}

	if len(latest) < len(sample) {
		t.Errorf("expected %d to be greater than %d", len(latest), len(sample))
	}
	//lastJobs := latest[len(latest)-len(sample):]

	//for k, v := range lastJobs {
	//	ks := len(sample) - k - 1
	//	t.Errorf("%d %d", k, ks)
	//	eqSample := sample[ks]
	//	if v.Title != eqSample.Title {
	//		t.Errorf("expected %s got %s", eqSample.Title, v.Title)
	//	}
	//}

}

// TestJobQuery is a test suite for all functions which interact with database that
// are dealing with the User model.
func TestUserQuery(t *testing.T) {
	sample := []struct {
		name, email, pass string
	}{
		{"gernest", "gernest@zedlist.io", "mypass"},
		{"zedlist", "zedlist@zedlist.io", "myscarypass"},
	}
	users := []*models.User{}
	for _, v := range sample {
		f := forms.Register{}
		f.FirstName = v.name
		f.MiddleName = v.name
		f.LastName = v.name
		f.Email = v.email
		f.Password = v.pass

		// CreateNewUser
		usr, err := CreateNewUser(f)
		if err != nil {
			t.Errorf("creating new user %v", err)
		}
		users = append(users, usr)
	}
	for _, v := range users {

		// GetUserByID
		usr, err := GetUserByID(v.ID)
		if err != nil {
			t.Errorf("getting user by id %v", err)
		}
		if usr.ID != v.ID {
			t.Errorf("expected %d got %d", v.ID, usr.ID)
		}

		// GetUserByEmail
		eUsr, err := GetUserByEmail(v.Email)
		if err != nil {
			t.Errorf("getting user by email %v", err)
		}
		if eUsr.Email != v.Email {
			t.Errorf("epected %s got %s", v.Email, eUsr.Email)
		}
	}

	// Failure cases

	_, err := GetUserByID(2000)
	if err == nil {
		t.Error("expected error got nil instead")
	}
	_, err = GetUserByEmail("bogus")
	if err == nil {
		t.Error("expected error got nilinstead")
	}

	//
	// AuthenticateUserByEmail
	//
	loginForm := forms.Login{
		Email:    sample[0].email,
		Password: sample[0].pass,
	}

	// Passing case
	usr, err := AuthenticateUserByEmail(loginForm)
	if err != nil {
		t.Errorf("authenticating user by email %v", err)
	}
	if usr.ID != users[0].ID {
		t.Errorf("expected %d got %d", users[0].ID, usr.ID)
	}

	// Wrong email
	loginForm.Email = "bogue"
	_, err = AuthenticateUserByEmail(loginForm)
	if err == nil {
		t.Error("expected error got nil instead")
	}

	// Wrong password
	loginForm.Email = sample[0].email
	loginForm.Password = "Ohmygawd"
	_, err = AuthenticateUserByEmail(loginForm)
	if err == nil {
		t.Error("expected error got nil instead")
	}
}

func TestPersonQuery(t *testing.T) {
	err := SampleUser()
	if err != nil {
		t.Errorf("creating sample user %v", err)
	}
	sampleUser, err := GetUserByEmail("root@home.com")
	if err != nil {
		t.Errorf("getting sample user %v", err)
	}

	if sampleUser == nil {
		t.Fatal("expected sample user got nil")
	}

	//
	//	GetPersonByUserID
	//
	person, err := GetPersonByUserID(sampleUser.ID)
	if err != nil {
		t.Errorf("getting person %v", err)
	}

	// check if the person name is also loaded
	if person.PersonName.GivenName != "root" {
		t.Errorf("expected root got %s", person.PersonName.GivenName)
	}

	//
	//	PersonCreateJob
	//
	jobForm := forms.JobForm{Title: "whacko job"}
	err = PersonCreateJob(person, jobForm)
	if err != nil {
		t.Errorf("creating job %v", err)
	}
}

func TestResumeQuery(t *testing.T) {
	p := &models.Person{
		AboutMe: "rocket scientist",
	}
	err := Create(p)
	if err != nil {
		t.Error(err)
	}

	// Add a sample resumes
	for _, v := range []string{"one", "two", "three"} {
		resume := models.SampleResume()
		resume.Name = v

		//
		//	CreateResume
		//
		rerr := CreateResume(p, resume)
		if rerr != nil {
			t.Errorf("creating resume %v", rerr)
		}
	}

	//
	//	GetResumeByID
	//
	resume, err := GetResumeByID(p.Resumes[0].ID)
	if err != nil {
		t.Errorf("getting resume %v", err)
	}

	// check whether the ResumeBasic was loaded
	if resume.ResumeBasic.Name != "John Doe" {
		t.Errorf("expected Jon Doe got %s", resume.ResumeBasic.Name)
	}

	//
	//	GetAllPersonResumes
	//
	resumes, err := GetAllPersonResumes(p)
	if err != nil {
		t.Errorf("getting all erson resumes %v", err)
	}
	if len(resumes) != 3 {
		t.Errorf("expected 3 got %d instead", len(resumes))
	}
}
