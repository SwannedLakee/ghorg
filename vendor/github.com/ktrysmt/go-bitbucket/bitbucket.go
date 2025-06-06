package bitbucket

import "context"

type users interface {
	Get(username string) (*User, error)
	Followers(username string) (interface{}, error)
	Following(username string) (interface{}, error)
	Repositories(username string) (interface{}, error)
}

type user interface {
	Profile() (*User, error)
	Emails() (interface{}, error)
}

type pullrequests interface {
	Create(opt PullRequestsOptions) (interface{}, error)
	Update(opt PullRequestsOptions) (interface{}, error)
	List(opt PullRequestsOptions) (interface{}, error)
	Get(opt PullRequestsOptions) (interface{}, error)
	Activities(opt PullRequestsOptions) (interface{}, error)
	Activity(opt PullRequestsOptions) (interface{}, error)
	Commits(opt PullRequestsOptions) (interface{}, error)
	Patch(opt PullRequestsOptions) (interface{}, error)
	Diff(opt PullRequestsOptions) (interface{}, error)
	Merge(opt PullRequestsOptions) (interface{}, error)
	Decline(opt PullRequestsOptions) (interface{}, error)
}
type workspace interface {
	GetProject(opt ProjectOptions) (*Project, error)
	CreateProject(opt ProjectOptions) (*Project, error)
}

type issues interface {
	Gets(io *IssuesOptions) (interface{}, error)
	Get(io *IssuesOptions) (interface{}, error)
	Delete(io *IssuesOptions) (interface{}, error)
	Update(io *IssuesOptions) (interface{}, error)
	Create(io *IssuesOptions) (interface{}, error)
	GetVote(io *IssuesOptions) (bool, interface{}, error)
	PutVote(io *IssuesOptions) error
	DeleteVote(io *IssuesOptions) error
	GetWatch(io *IssuesOptions) (bool, interface{}, error)
	PutWatch(io *IssuesOptions) error
	DeleteWatch(io *IssuesOptions) error
	GetComments(ico *IssueCommentsOptions) (interface{}, error)
	CreateComment(ico *IssueCommentsOptions) (interface{}, error)
	GetComment(ico *IssueCommentsOptions) (interface{}, error)
	UpdateComment(ico *IssueCommentsOptions) (interface{}, error)
	DeleteComment(ico *IssueCommentsOptions) (interface{}, error)
	GetChanges(ico *IssueChangesOptions) (interface{}, error)
	CreateChange(ico *IssueChangesOptions) (interface{}, error)
	GetChange(ico *IssueChangesOptions) (interface{}, error)
}

type repository interface {
	Get(opt RepositoryOptions) (*Repository, error)
	Create(opt RepositoryOptions) (*Repository, error)
	Delete(opt RepositoryOptions) (interface{}, error)
	ListWatchers(opt RepositoryOptions) (interface{}, error)
	ListForks(opt RepositoryOptions) (interface{}, error)
	ListDefaultReviewers(opt RepositoryOptions) (*DefaultReviewers, error)
	GetDefaultReviewer(opt RepositoryDefaultReviewerOptions) (*DefaultReviewer, error)
	AddDefaultReviewer(opt RepositoryDefaultReviewerOptions) (*DefaultReviewer, error)
	DeleteDefaultReviewer(opt RepositoryDefaultReviewerOptions) (interface{}, error)
	UpdatePipelineConfig(opt RepositoryPipelineOptions) (*Pipeline, error)
	ListPipelineVariables(opt RepositoryPipelineVariablesOptions) (*PipelineVariables, error)
	AddPipelineVariable(opt RepositoryPipelineVariableOptions) (*PipelineVariable, error)
	DeletePipelineVariable(opt RepositoryPipelineVariableDeleteOptions) (interface{}, error)
	AddPipelineKeyPair(opt RepositoryPipelineKeyPairOptions) (*PipelineKeyPair, error)
	UpdatePipelineBuildNumber(opt RepositoryPipelineBuildNumberOptions) (*PipelineBuildNumber, error)
	ListFiles(opt RepositoryFilesOptions) (*[]RepositoryFile, error)
	GetFileBlob(opt RepositoryBlobOptions) (*RepositoryBlob, error)
	ListBranches(opt RepositoryBranchOptions) (*RepositoryBranches, error)
	BranchingModel(opt RepositoryBranchingModelOptions) (*BranchingModel, error)
	ListEnvironments(opt RepositoryEnvironmentsOptions) (*Environments, error)
	AddEnvironment(opt RepositoryEnvironmentOptions) (*Environment, error)
	DeleteEnvironment(opt RepositoryEnvironmentDeleteOptions) (interface{}, error)
	GetEnvironment(opt RepositoryEnvironmentOptions) (*Environment, error)
	ListDeploymentVariables(opt RepositoryDeploymentVariablesOptions) (*DeploymentVariables, error)
	AddDeploymentVariable(opt RepositoryDeploymentVariableOptions) (*DeploymentVariable, error)
	DeleteDeploymentVariable(opt RepositoryDeploymentVariableDeleteOptions) (interface{}, error)
	UpdateDeploymentVariable(opt RepositoryDeploymentVariableOptions) (*DeploymentVariable, error)
}

type repositories interface {
	ListForAccount(opt RepositoriesOptions) (interface{}, error)
	ListForTeam(opt RepositoriesOptions) (interface{}, error)
	ListProject(opt RepositoriesOptions) (interface{}, error)
	ListPublic() (interface{}, error)
}

type commits interface {
	GetCommits(opt CommitsOptions) (interface{}, error)
	GetCommit(opt CommitsOptions) (interface{}, error)
	GetCommitComments(opt CommitsOptions) (interface{}, error)
	GetCommitComment(opt CommitsOptions) (interface{}, error)
	GetCommitStatus(opt CommitsOptions) (interface{}, error)
	GiveApprove(opt CommitsOptions) (interface{}, error)
	RemoveApprove(opt CommitsOptions) (interface{}, error)
	CreateCommitStatus(cmo CommitsOptions, cso CommitStatusOptions) (interface{}, error)
}

type branchrestrictions interface {
	Gets(opt BranchRestrictionsOptions) (interface{}, error)
	Get(opt BranchRestrictionsOptions) (interface{}, error)
	Create(opt BranchRestrictionsOptions) (interface{}, error)
	Update(opt BranchRestrictionsOptions) (interface{}, error)
	Delete(opt BranchRestrictionsOptions) (interface{}, error)
}

type diff interface {
	GetDiff(opt DiffOptions) (interface{}, error)
	GetPatch(opt DiffOptions) (interface{}, error)
}

type webhooks interface {
	Gets(opt WebhooksOptions) (interface{}, error)
	Get(opt WebhooksOptions) (interface{}, error)
	Create(opt WebhooksOptions) (interface{}, error)
	Update(opt WebhooksOptions) (interface{}, error)
	Delete(opt WebhooksOptions) (interface{}, error)
}

type teams interface {
	List(role string) (interface{}, error) // [WIP?] role=[admin|contributor|member]
	Profile(teamname string) (interface{}, error)
	Members(teamname string) (interface{}, error)
	Followers(teamname string) (interface{}, error)
	Following(teamname string) (interface{}, error)
	Repositories(teamname string) (interface{}, error)
	Projects(teamname string) (interface{}, error)
}

type pipelines interface {
	List(po *PipelinesOptions) (interface{}, error)
	Get(po *PipelinesOptions) (interface{}, error)
	ListSteps(po *PipelinesOptions) (interface{}, error)
	GetStep(po *PipelinesOptions) (interface{}, error)
	GetLog(po *PipelinesOptions) (string, error)
}

type RepositoriesOptions struct {
	Owner   string  `json:"owner"`
	Project string  `json:"project"`
	Role    string  `json:"role"` // role=[owner|admin|contributor|member]
	Page    *int    `json:"page"`
	Keyword *string `json:"keyword"`
}

type RepositoryOptions struct {
	Uuid     string `json:"uuid"`
	Owner    string `json:"owner"`
	RepoSlug string `json:"repo_slug"`
	Scm      string `json:"scm"`
	//	Name        string `json:"name"`
	IsPrivate   string `json:"is_private"`
	Description string `json:"description"`
	ForkPolicy  string `json:"fork_policy"`
	Language    string `json:"language"`
	HasIssues   string `json:"has_issues"`
	HasWiki     string `json:"has_wiki"`
	Project     string `json:"project"`
	ctx         context.Context
}

func (ro *RepositoryOptions) WithContext(ctx context.Context) *RepositoryOptions {
	ro.ctx = ctx
	return ro
}

type RepositoryForkOptions struct {
	FromOwner string `json:"from_owner"`
	FromSlug  string `json:"from_slug"`
	Owner     string `json:"owner"`
	// TODO: does the API supports specifying  slug on forks?
	// see: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Bworkspace%7D/%7Brepo_slug%7D/forks#post
	Name        string `json:"name"`
	IsPrivate   string `json:"is_private"`
	Description string `json:"description"`
	ForkPolicy  string `json:"fork_policy"`
	Language    string `json:"language"`
	HasIssues   string `json:"has_issues"`
	HasWiki     string `json:"has_wiki"`
	Project     string `json:"project"`
	ctx         context.Context
}

func (fo *RepositoryForkOptions) WithContext(ctx context.Context) *RepositoryForkOptions {
	fo.ctx = ctx
	return fo
}

type RepositoryFilesOptions struct {
	Owner    string `json:"owner"`
	RepoSlug string `json:"repo_slug"`
	Ref      string `json:"ref"`
	Path     string `json:"path"`
	MaxDepth int    `json:"max_depth"`
}

type RepositoryBlobOptions struct {
	Owner    string `json:"owner"`
	RepoSlug string `json:"repo_slug"`
	Ref      string `json:"ref"`
	Path     string `json:"path"`
}

type File struct {
	Path string
	Name string
}

// Based on https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Bworkspace%7D/%7Brepo_slug%7D/src#post
type RepositoryBlobWriteOptions struct {
	Owner         string   `json:"owner"`
	RepoSlug      string   `json:"repo_slug"`
	FilePath      string   `json:"filepath"`
	FileName      string   `json:"filename"`
	Files         []File   `json:"files"`
	FilesToDelete []string `json:"files_to_delete"`
	Author        string   `json:"author"`
	Message       string   `json:"message"`
	Branch        string   `json:"branch"`
	ctx           context.Context
}

func (ro *RepositoryBlobWriteOptions) WithContext(ctx context.Context) *RepositoryBlobWriteOptions {
	ro.ctx = ctx
	return ro
}

// RepositoryRefOptions represents the options for describing a repository's refs (i.e.
// tags and branches). The field BranchFlg is a boolean that is indicates whether a specific
// RepositoryRefOptions instance is meant for Branch specific set of api methods.
type RepositoryRefOptions struct {
	Owner     string `json:"owner"`
	RepoSlug  string `json:"repo_slug"`
	Query     string `json:"query"`
	Sort      string `json:"sort"`
	PageNum   int    `json:"page"`
	Pagelen   int    `json:"pagelen"`
	MaxDepth  int    `json:"max_depth"`
	Name      string `json:"name"`
	BranchFlg bool
}

type RepositoryBranchOptions struct {
	Owner      string `json:"owner"`
	RepoSlug   string `json:"repo_slug"`
	Query      string `json:"query"`
	Sort       string `json:"sort"`
	PageNum    int    `json:"page"`
	Pagelen    int    `json:"pagelen"`
	MaxDepth   int    `json:"max_depth"`
	BranchName string `json:"branch_name"`
}

type RepositoryBranchCreationOptions struct {
	Owner    string                 `json:"owner"`
	RepoSlug string                 `json:"repo_slug"`
	Name     string                 `json:"name"`
	Target   RepositoryBranchTarget `json:"target"`
}

type RepositoryBranchDeleteOptions struct {
	Owner    string `json:"owner"`
	RepoSlug string `json:"repo_slug"`
	RepoUUID string `json:"uuid"`
	RefName  string `json:"name"`
	RefUUID  string `json:"uuid"`
}

type RepositoryBranchTarget struct {
	Hash string `json:"hash"`
}

type RepositoryTagOptions struct {
	Owner    string `json:"owner"`
	RepoSlug string `json:"repo_slug"`
	Query    string `json:"q"`
	Sort     string `json:"sort"`
	PageNum  int    `json:"page"`
	Pagelen  int    `json:"pagelen"`
	MaxDepth int    `json:"max_depth"`
}

type RepositoryTagCreationOptions struct {
	Owner    string              `json:"owner"`
	RepoSlug string              `json:"repo_slug"`
	Name     string              `json:"name"`
	Target   RepositoryTagTarget `json:"target"`
}

type RepositoryTagTarget struct {
	Hash string `json:"hash"`
}

type PullRequestsOptions struct {
	ID                string   `json:"id"`
	CommentID         string   `json:"comment_id"`
	Owner             string   `json:"owner"`
	RepoSlug          string   `json:"repo_slug"`
	Title             string   `json:"title"`
	Description       string   `json:"description"`
	CloseSourceBranch bool     `json:"close_source_branch"`
	SourceBranch      string   `json:"source_branch"`
	SourceRepository  string   `json:"source_repository"`
	DestinationBranch string   `json:"destination_branch"`
	DestinationCommit string   `json:"destination_repository"`
	Message           string   `json:"message"`
	Reviewers         []string `json:"reviewers"`
	States            []string `json:"states"`
	Query             string   `json:"query"`
	Sort              string   `json:"sort"`
	Draft             bool     `json:"draft"`
	Commit            string   `json:"commit"`
	ctx               context.Context
}

func (po *PullRequestsOptions) WithContext(ctx context.Context) *PullRequestsOptions {
	po.ctx = ctx
	return po
}

type PullRequestCommentOptions struct {
	Owner         string `json:"owner"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"id"`
	Content       string `json:"content"`
	CommentId     string `json:"-"`
	Parent        *int   `json:"parent"`
	ctx           context.Context
}

func (pco *PullRequestCommentOptions) WithContext(ctx context.Context) *PullRequestCommentOptions {
	pco.ctx = ctx
	return pco
}

type IssuesOptions struct {
	ID        string   `json:"id"`
	Owner     string   `json:"owner"`
	RepoSlug  string   `json:"repo_slug"`
	States    []string `json:"states"`
	Query     string   `json:"query"`
	Sort      string   `json:"sort"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	State     string   `json:"state"`
	Kind      string   `json:"kind"`
	Milestone string   `json:"milestone"`
	Component string   `json:"component"`
	Priority  string   `json:"priority"`
	Version   string   `json:"version"`
	Assignee  string   `json:"assignee"`
	ctx       context.Context
}

func (io *IssuesOptions) WithContext(ctx context.Context) *IssuesOptions {
	io.ctx = ctx
	return io
}

type IssueCommentsOptions struct {
	IssuesOptions
	Query          string `json:"query"`
	Sort           string `json:"sort"`
	CommentContent string `json:"comment_content"`
	CommentID      string `json:"comment_id"`
}

type IssueChangesOptions struct {
	IssuesOptions
	Query    string `json:"query"`
	Sort     string `json:"sort"`
	Message  string `json:"message"`
	ChangeID string `json:"change_id"`
	Changes  []struct {
		Type     string
		NewValue string
	} `json:"changes"`
}

type CommitsOptions struct {
	Owner       string `json:"owner"`
	RepoSlug    string `json:"repo_slug"`
	Revision    string `json:"revision"`
	Branchortag string `json:"branchortag"`
	Include     string `json:"include"`
	Exclude     string `json:"exclude"`
	CommentID   string `json:"comment_id"`
	Page        *int   `json:"page"`
	ctx         context.Context
}

func (cm *CommitsOptions) WithContext(ctx context.Context) *CommitsOptions {
	cm.ctx = ctx
	return cm
}

type CommitStatusOptions struct {
	Key         string `json:"key"`
	Url         string `json:"url"`
	State       string `json:"state"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BranchRestrictionsOptions struct {
	Owner    string            `json:"owner"`
	RepoSlug string            `json:"repo_slug"`
	ID       string            `json:"id"`
	Groups   map[string]string `json:"groups"`
	Pattern  string            `json:"pattern"`
	Users    []string          `json:"users"`
	Kind     string            `json:"kind"`
	FullSlug string            `json:"full_slug"`
	Name     string            `json:"name"`
	Value    interface{}       `json:"value"`
	ctx      context.Context
}

func (b *BranchRestrictionsOptions) WithContext(ctx context.Context) *BranchRestrictionsOptions {
	b.ctx = ctx
	return b
}

type DiffOptions struct {
	Owner             string `json:"owner"`
	RepoSlug          string `json:"repo_slug"`
	Spec              string `json:"spec"`
	Context           int    `json:"context"`
	Path              string `json:"path"`
	FromPullRequestID int    `json:"from_pullrequest_id"`
	Whitespace        bool   `json:"ignore_whitespace"`
	Binary            bool   `json:"binary"`
	Renames           bool   `json:"renames"`
	Topic             bool   `json:"topic"`
}

type DiffStatOptions struct {
	Owner             string `json:"owner"`
	RepoSlug          string `json:"repo_slug"`
	Spec              string `json:"spec"`
	FromPullRequestID int    `json:"from_pullrequest_id"`
	Whitespace        bool   `json:"ignore_whitespace"`
	// Deprecated: Merge is deprecated use Topic
	Merge    bool   `json:"merge"`
	Path     string `json:"path"`
	Renames  bool   `json:"renames"`
	Topic    bool   `json:"topic"`
	PageNum  int    `json:"page"`
	Pagelen  int    `json:"pagelen"`
	MaxDepth int    `json:"max_depth"`
	Fields   []string
}

type WebhooksOptions struct {
	Owner       string   `json:"owner"`
	RepoSlug    string   `json:"repo_slug"`
	Uuid        string   `json:"uuid"`
	Secret      string   `json:"secret"`
	Description string   `json:"description"`
	Url         string   `json:"url"`
	Active      bool     `json:"active"`
	Events      []string `json:"events"` // EX: {'repo:push','issue:created',..} REF: https://bit.ly/3FjRHHu
	ctx         context.Context
}

func (wo *WebhooksOptions) WithContext(ctx context.Context) *WebhooksOptions {
	wo.ctx = ctx
	return wo
}

type RepositoryPipelineOptions struct {
	Owner    string `json:"owner"`
	RepoSlug string `json:"repo_slug"`
	Enabled  bool   `json:"has_pipelines"`
}

type RepositoryDefaultReviewerOptions struct {
	Owner    string `json:"owner"`
	RepoSlug string `json:"repo_slug"`
	Username string `json:"username"`
}

type RepositoryPipelineVariablesOptions struct {
	Owner    string `json:"owner"`
	RepoSlug string `json:"repo_slug"`
	Query    string `json:"q"`
	Sort     string `json:"sort"`
	PageNum  int    `json:"page"`
	Pagelen  int    `json:"pagelen"`
	MaxDepth int    `json:"max_depth"`
}

type RepositoryPipelineVariableOptions struct {
	Owner    string `json:"owner"`
	RepoSlug string `json:"repo_slug"`
	Uuid     string `json:"uuid"`
	Key      string `json:"key"`
	Value    string `json:"value"`
	Secured  bool   `json:"secured"`
	ctx      context.Context
}

func (rpvo *RepositoryPipelineVariableOptions) WithContext(ctx context.Context) *RepositoryPipelineVariableOptions {
	rpvo.ctx = ctx
	return rpvo
}

type RepositoryPipelineVariableDeleteOptions struct {
	Owner    string `json:"owner"`
	RepoSlug string `json:"repo_slug"`
	Uuid     string `json:"uuid"`
}

type RepositoryPipelineKeyPairOptions struct {
	Owner      string `json:"owner"`
	RepoSlug   string `json:"repo_slug"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

type RepositoryPipelineBuildNumberOptions struct {
	Owner    string `json:"owner"`
	RepoSlug string `json:"repo_slug"`
	Next     int    `json:"next"`
}

type RepositoryBranchingModelOptions struct {
	Owner    string `json:"owner"`
	RepoSlug string `json:"repo_slug"`
}

type DownloadsOptions struct {
	Owner    string `json:"owner"`
	RepoSlug string `json:"repo_slug"`
	FilePath string `json:"filepath"`
	FileName string `json:"filename"`
	Files    []File `json:"files"`
	ctx      context.Context
}

func (do *DownloadsOptions) WithContext(ctx context.Context) *DownloadsOptions {
	do.ctx = ctx
	return do
}

type PageRes struct {
	Page     int32 `json:"page"`
	PageLen  int32 `json:"pagelen"`
	MaxDepth int32 `json:"max_depth"`
	Size     int32 `json:"size"`
}

type PipelinesOptions struct {
	Owner    string `json:"owner"`
	Page     int    `json:"page"`
	RepoSlug string `json:"repo_slug"`
	Query    string `json:"query"`
	Sort     string `json:"sort"`
	IDOrUuid string `json:"ID"`
	StepUuid string `json:"StepUUID"`
}

type RepositoryEnvironmentsOptions struct {
	Owner    string `json:"owner"`
	RepoSlug string `json:"repo_slug"`
}

type RepositoryGroupPermissionsOptions struct {
	Owner      string `json:"owner"`
	RepoSlug   string `json:"repo_slug"`
	Group      string `json:"group"`
	Permission string `json:"permission"`
}

type RepositoryUserPermissionsOptions struct {
	Owner      string `json:"owner"`
	RepoSlug   string `json:"repo_slug"`
	User       string `json:"user"`
	Permission string `json:"permission"`
}

type RepositoryEnvironmentTypeOption int

const (
	Test RepositoryEnvironmentTypeOption = iota
	Staging
	Production
)

func (e RepositoryEnvironmentTypeOption) String() string {
	return [...]string{"Test", "Staging", "Production"}[e]
}

type RepositoryEnvironmentOptions struct {
	Owner           string                          `json:"owner"`
	RepoSlug        string                          `json:"repo_slug"`
	Uuid            string                          `json:"uuid"`
	Name            string                          `json:"name"`
	EnvironmentType RepositoryEnvironmentTypeOption `json:"environment_type"`
	Rank            int                             `json:"rank"`
	ctx             context.Context
}

func (reo *RepositoryEnvironmentOptions) WithContext(ctx context.Context) *RepositoryEnvironmentOptions {
	reo.ctx = ctx
	return reo
}

type RepositoryEnvironmentDeleteOptions struct {
	Owner    string `json:"owner"`
	RepoSlug string `json:"repo_slug"`
	Uuid     string `json:"uuid"`
}

type RepositoryDeploymentVariablesOptions struct {
	Owner       string       `json:"owner"`
	RepoSlug    string       `json:"repo_slug"`
	Environment *Environment `json:"environment"`
	Query       string       `json:"q"`
	Sort        string       `json:"sort"`
	PageNum     int          `json:"page"`
	Pagelen     int          `json:"pagelen"`
	MaxDepth    int          `json:"max_depth"`
}

type RepositoryDeploymentVariableOptions struct {
	Owner       string       `json:"owner"`
	RepoSlug    string       `json:"repo_slug"`
	Environment *Environment `json:"environment"`
	Uuid        string       `json:"uuid"`
	Key         string       `json:"key"`
	Value       string       `json:"value"`
	Secured     bool         `json:"secured"`
	ctx         context.Context
}

func (rdvo *RepositoryDeploymentVariableOptions) WithContext(ctx context.Context) *RepositoryDeploymentVariableOptions {
	rdvo.ctx = ctx
	return rdvo
}

type RepositoryDeploymentVariableDeleteOptions struct {
	Owner       string       `json:"owner"`
	RepoSlug    string       `json:"repo_slug"`
	Environment *Environment `json:"environment"`
	Uuid        string       `json:"uuid"`
}

type DeployKeyOptions struct {
	Owner    string `json:"owner"`
	RepoSlug string `json:"repo_slug"`
	Id       int    `json:"id"`
	Label    string `json:"label"`
	Key      string `json:"key"`
	ctx      context.Context
}

func (dk *DeployKeyOptions) WithContext(ctx context.Context) *DeployKeyOptions {
	dk.ctx = ctx
	return dk
}

type SSHKeyOptions struct {
	Owner string `json:"owner"`
	Uuid  string `json:"uuid"`
	Label string `json:"label"`
	Key   string `json:"key"`
}
